package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestItemSaver(t *testing.T) {
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1737015172",
		Type: "zhenai",
		Id:   "1737015172",
		Payload: model.Profile{
			Name: "name",
			Car:  "car",
		},
	}
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.245.137:9200"),
		// Must turn off sniff in docker ,用来维护集群状态的
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, "dating_profile", item)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type(item.Type).Id(item.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

}
