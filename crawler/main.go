package main

import (
	"./engine"
	"./parser/zhenai"
	"./scheduler"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
}
