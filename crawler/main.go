package main

import (
	"./persist"
	"./engine"
	"./model"
	"./scheduler"
	"./parser/zhenai"
	"fmt"
)

const url = "http://www.zhenai.com/zhenghun"
const shUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	run()
	//saveToElasticsearch()
}

func run() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(),
	}
	//e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e.Run(engine.Request{Url: shUrl, ParserFunc: parser.ParseProfileList})
}

func saveToElasticsearch() {
	profile := model.Profile{
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
	}
	id, err := persist.Save(profile)
	fmt.Println(id, err)
}
