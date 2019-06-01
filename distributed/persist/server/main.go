package main

import (
	"crawler/distributed/persist"
	rpc "crawler/distributed/rpcSupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serverRPC(":1234", "dating_profile2"))
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
