package parser

import (
	"../../engine"
	"../../model"
	"regexp"
)

const profileListRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`


//城市个人信息列表解析
func ParseProfileList(body []byte) engine.ParseResult {
	compile := regexp.MustCompile(profileListRe)
	profileSubMatch := compile.FindAllStringSubmatch(string(body), -1)
	parseResult := engine.ParseResult{}
	for _, item := range profileSubMatch {
		var profile model.Profile
		profile.Name = item[2]
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: item[1],
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, profile)
			},
		})
		parseResult.Items = append(parseResult.Items, profile.Name)
	}
	return parseResult
}
