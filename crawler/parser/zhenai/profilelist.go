package parser

import (
	"../../engine"
	"../../model"
	"fmt"
	"regexp"
)

const profileListRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const profileNextRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z/]+)">下一页</a>`

//城市个人信息列表解析
func ParseProfileList(body []byte) engine.ParseResult {
	body2 := string(body)
	compile := regexp.MustCompile(profileListRe)
	profileSubMatch := compile.FindAllStringSubmatch(body2, -1)
	fmt.Println("profileSubMatch",profileSubMatch)
	parseResult := engine.ParseResult{}
	for _, item := range profileSubMatch {
		addDataRequestQ(item[2], item[1], &parseResult)
	}
	nextPageSubMatch := regexp.MustCompile(profileNextRe).FindAllStringSubmatch(body2, -1)
	fmt.Println("nextPageSubMatch",nextPageSubMatch)
	/*for _, item := range nextPageSubMatch {
		var profile model.Profile
		//profile.Name = name
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: item[1],
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseCityList(bytes)
			},
		})
		parseResult.Items = append(parseResult.Items, profile.Name)
	}*/

	return parseResult
}

func addDataRequestQ(name string, url string, parseResult *engine.ParseResult) {
	var profile model.Profile
	profile.Name = name
	parseResult.Requests = append(parseResult.Requests, engine.Request{
		Url: url,
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return ParseProfile(bytes, profile)
		},
	})
	parseResult.Items = append(parseResult.Items, profile.Name)
}
