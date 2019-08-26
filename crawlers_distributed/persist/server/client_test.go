package main

import (
	"../../../crawler/engine"
	"../../../crawler/model"
	"../../rpcsupport"
	"fmt"
	"testing"
	".."
	"time"
)

const host = ":1234"

func TestItemSaver(t *testing.T) {

	go persist.StartElasticServer(host, "test1")
	time.Sleep(time.Second * 5)
	sendRequest()
}

func sendRequest() {

	item := engine.Item{
		Id:   "test_id",
		Url:  "http://www.baidu3.com",
		Type: "zhenai",
		Payload: model.Profile{
			Age:        34,
			Height:     166,
			Weight:     61,
			Income:     "2000-5000",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "水平",
			Occupation: "人事/行政",
			Marriage:   "未婚",
			House:      "已购房",
			HuKou:      "上海",
			Education:  "大学本科",
			Car:        "已购车",
		},
	}
	var rpcClient rpcsupport.AppRpcClient
	err := rpcClient.NewRpcClient(host)
	if err != nil {
		panic(err)
	}
	var result2 string
	result2, err = rpcClient.CallFun("ItemSavesService.Save", item)
	fmt.Println(result2, err)
}
