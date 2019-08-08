package persist

import (
	"context"
	"github.com/olivere/elastic"
	"fmt"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount ++
			fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			/*_, err := Save(item)
			if err != nil {
				log.Error("Item Save: error item %v\n%v", item, err)
			}*/
		}
	}()
	return out
}

func Save(item interface{}) (id string, err error) {
	index := "dating_profile"
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return "", err
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index) /*.BodyJson(item)*/ .Do(context.Background())
		if err != nil {
			// Handle error
			return "", err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			//panic(err)
		}
	}

	response, err := client.Index().Index(index).
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	//fmt.Printf("%+v\n", response)
	return response.Id, nil

}
