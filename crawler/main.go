package main

import (
	"./engine"
	"./scheduler"
	"./parser/zhenai"
	"./persist"
)

const url = "http://www.zhenai.com/zhenghun"
const shUrl = "http://www.zhenai.com/zhenghun/shanghai"

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(),
	}
	//e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e.Run(engine.Request{Url: shUrl, ParserFunc: parser.ParseProfileList})

}
