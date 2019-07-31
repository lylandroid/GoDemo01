package main

import (
	"./engine"
	"./parser/zhenai"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
}
