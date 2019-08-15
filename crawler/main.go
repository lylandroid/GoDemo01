package main

import (
	"./engine"
	"./parser/zhenai"
	"./persist"
	"./scheduler"
	"./model"
)

const url = "http://www.zhenai.com/zhenghun"
const shUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	run()
	//saveToElasticsearch()
	//inseartTestData()
}

var index = "dating_profile"

func run() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(index),
	}
	//e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	//e.Run(engine.Request{Url: shUrl, ParserFunc: parser.ParseProfileList})
}

func inseartTestData()  {
	client, _ := persist.NewClient()
	persist.Save(index, client, engine.Item{
		Id:   "test_id",
		Url:  "http://www.baidu.com",
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
			/*Car:        "已购车",*/
		},
	})
}

/*func saveToElasticsearch() {
	url := ""
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
	item := engine.Item{
		Url:     url,
		_Id:      parser.ExtractId(url),
		Payload: profile,
	}

	err := persist.Save(item)
	fmt.Println(id, err)
}*/
