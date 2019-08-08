package persist

import (
	"../engine"
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/olivere/elastic"
	"../model"
	"fmt"
	"github.com/pkg/errors"
)

func ItemServer() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		itemCount2 := 0
		for {
			item := <-out
			itemCount ++
			//fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			switch item.Payload.(type) {
			case string:
				fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			case model.Profile:
				itemCount2++
				err := Save(item)
				if err != nil {
					log.Error("Item Save: error item %v\t%v", item, err)
				} else {
					fmt.Printf("Item Saves Success: %d %v\n", itemCount2, item)
				}
			}

		}
	}()
	return out
}

func Save(item engine.Item) error {
	index := "dating_profile"
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index) /*.BodyJson(item)*/ .Do(context.Background())
		if err != nil {
			// Handle error
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			//panic(err)
		}
	}
	if item.Type == "" {
		return errors.New("most supply Type not null")
	}
	indexServer := client.Index().Index(index)
	if item.Id != "" {
		indexServer.Id(item.Id)
	}
	_, err = indexServer.
		Type( /*"zhenai"*/ item.Type).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil

}
