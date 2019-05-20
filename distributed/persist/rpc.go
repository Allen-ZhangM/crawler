package persist

import (
	"crawler/engine"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (i *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(i.Client, i.Index, item)
	log.Printf("Item saver: got item : %v", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Item saver error : item : %v", item)
	}
	return err
}
