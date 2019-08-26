package client

import (
	"../../../crawler/engine"
	"../../../crawler/model"
	"../../rpcsupport"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"../../config"
)

func ItemServer(host string) (chan engine.Item, error) {
	appRpcClient := rpcsupport.AppRpcClient{}
	err := appRpcClient.NewRpcClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		//client, err := persist.NewElasticClient()
		//if err != nil {
		//	panic(err)
		//}
		itemCount := 0
		itemCount2 := 0
		for {
			item := <-out
			itemCount++
			//fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			switch item.Payload.(type) {
			case string:
				fmt.Printf("Item Saves: %d %v\n", itemCount, item)
			case model.Profile:
				itemCount2++
				result, err := appRpcClient.CallFun(config.ItemSaverRPCApi, item)
				//err := persist.Save(index, client, item)
				if err != nil {
					log.Error("Item Save: error item %v\t%v\t %s \n", item, err, result)
				} else {
					fmt.Printf("Item Saves Success: %d %v \t %s \n", itemCount2, item, result)
				}
			}

		}
	}()
	return out, nil
}
