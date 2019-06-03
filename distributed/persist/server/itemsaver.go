package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	rpc "crawler/distributed/rpcSupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRPC(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serverRPC(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.245.137:9200"),
		// Must turn off sniff in docker ,用来维护集群状态的
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	rpc.ServerRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	return nil
}
