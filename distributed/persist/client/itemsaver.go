package client

import (
	"crawler/engine"
	"crawler/rpc"
	"log"
)

func ItemSaver(host string) chan engine.Item {
	client, e := rpc.NewClient(host)
	if e != nil {
		log.Println("elastic.NewClient error ")
	}
	result := ""
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := client.Call("ItemSaverService.Save", item, &result)
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
