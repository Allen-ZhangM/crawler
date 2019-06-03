package client

import (
	"crawler/distributed/config"
	rpc "crawler/distributed/rpcSupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) chan engine.Item {
	client, e := rpc.NewClient(host)
	if e != nil {
		log.Printf("rpc.NewClient(%v) error:%v ", host, e)
	}
	result := ""
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("call ItemSaverService.Save error : result:%s, err:%v", result, err)
			}
			if err != nil {
				log.Printf("Item saver error : item #%d: %v", itemCount, item)
			}
		}
	}()
	return out
}
