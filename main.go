package main

import (
	"bytes"
	"crawler/engine"
	"crawler/fetcher"
	"crawler/scheduler"
	"crawler/zhenaiwang/parser"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	runConcurrent()
	//runRequest()
}

func runConcurrent() {
	e := engine.ConcurrentEngine{
		Scheduler:   &shceduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.RunConcurrentRequest(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

func runSimple() {
	engine.RunRequest(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
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
