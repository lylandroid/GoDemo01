package parser

import (
	"../../engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//城市列表解析
func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	match := compile.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}
	count := 5
	for _, v := range match {
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{Url: string(v[1]), ParserFunc: ParseProfileList})
		parseResult.Items = append(parseResult.Items, string(v[2]))
		count--
		if count < 0 {
			break
		}
	}
	return parseResult
}
