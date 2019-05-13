package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan engine.Item {
	client, e := elastic.NewClient(
		elastic.SetURL("http://192.168.245.137:9200"),
		// Must turn off sniff in docker ,用来维护集群状态的
		elastic.SetSniff(false))
	if e != nil {
		log.Println("elastic.NewClient error ")
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(client, item)
			if err != nil {
				log.Printf("Item saver error : item #%d: %v", itemCount, item)
			}
		}
	}()
	return out
}

func save(client *elastic.Client, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply type ")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, e := indexService.Do(context.Background())
	if e != nil {
		return e
	}

	return nil
}
