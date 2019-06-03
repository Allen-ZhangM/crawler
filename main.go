package main

import (
	"bytes"
	"crawler/distributed/config"
	itemsaver "crawler/distributed/persist/client"
	"crawler/distributed/rpcSupport"
	worker "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/fetcher"
	"crawler/scheduler"
	"crawler/zhenaiwang/parser"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"strings"
)

func main() {
	runConcurrent()
	//runRequest()
}

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker_hosts host (comma separated)")
)

func runConcurrent() {
	flag.Parse()

	itemChan := itemsaver.ItemSaver(*itemSaverHost)

	processor := worker.CreateProcessor(createClientPool(strings.Split(*workerHosts, ",")))

	e := engine.ConcurrentEngine{
		Scheduler:        &shceduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.RunConcurrentRequest(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts {
		c, e := rpcSupport.NewClient(host)
		if e != nil {
			log.Println(e)
		} else {
			clients = append(clients, c)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}

func runSimple() {
	engine.RunRequest(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

func runFetcher() {
	body, err := fetcher.Fetch("http://album.zhenai.com/u/1002992717")
	if err != nil {
		log.Printf("Fetcher: error fetching url: %v", err)
		return
	}
	fmt.Printf("%s", []byte(body))
}

func runRequest() {
	req, err := http.NewRequest("GET", "http://album.zhenai.com/u/1002992717", bytes.NewBuffer(nil))
	//req.Header.Set(ReqHeaderKey, ReqHeaderValue)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Fetcher: error fetching url: %v", err)
	}
	if resp != nil {
		defer func() {

			err = resp.Body.Close()

		}()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", []byte(body))
	}
}
