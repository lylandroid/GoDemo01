package persist

import (
	"../engine"
	"../model"
	"context"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/olivere/elastic"
)

func ItemServer(index string) chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		client, err := NewElasticClient()
		if err != nil {
			panic(err)
		}
		itemCount := 0
		itemCount2 := 0
		for {
			item := <-out
			itemCount++
			//fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			switch item.Payload.(type) {
			case string:
				fmt.Printf("Item Saves: itemCount=%d \t item=%v\n", itemCount, item)
			case model.Profile:
				itemCount2++
				err := Save(index, client, item)
				if err != nil {
					log.Error("Item Save: error itemCount2=%d item=%v \t %v", itemCount2, item, err)
				} else {
					fmt.Printf("Item Saves Success: itemCount2=%d \t item=%v\n", itemCount2, item)
				}
			}

		}
	}()
	return out
}

func Save(index string, client *elastic.Client, item engine.Item) error {
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := client.CreateIndex(index) /*.BodyJson(&item)*/ .Do(context.Background())
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

func NewElasticClient() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetSniff(false))
}
