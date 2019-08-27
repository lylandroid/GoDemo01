package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
//17-6 13:05
//城市列表解析
func ParseCityList(_ string, body []byte, _ interface{}) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	match := compile.FindAllSubmatch(body, -1)
	parseResult := engine.ParseResult{}
	for _, v := range match {
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{Url: string(v[1]), Parser: engine.NewFuncParser(ParseProfileList,"ParseProfileList")})
		var url = string(v[1])
		parseResult.Items = append(parseResult.Items, /* string(v[2])*/ engine.Item{
			Url:     url,
			Id:      ExtractId(url),
			Payload: string(v[2]),
		})
	}
	return parseResult
}
