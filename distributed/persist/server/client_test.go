package main

import (
	"crawler/distributed/config"
	rpc "crawler/distributed/rpcSupport"
	"crawler/engine"
	"crawler/model"
	"log"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serverRPC(host, "test1")
	time.Sleep(time.Second)

	client, _ := rpc.NewClient(host)

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1737015172",
		Type: "zhenai",
		Id:   "1737015172",
		Payload: model.Profile{
			Name: "name",
			Car:  "car",
		},
	}

	result := ""
	err := client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		log.Printf("call ItemSaverService.Save error : result:%s, err:%v", result, err)
	}

}
