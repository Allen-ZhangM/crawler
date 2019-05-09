package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++
			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	client, e := elastic.NewClient(
		elastic.SetURL("http://192.168.245.137:9200"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	response, e := client.Index().Index("dating_profile").Type("zhenai").
		BodyJson(item).Do(context.Background())
	if e != nil {
		panic(e)
	}

	fmt.Println(response)
}
