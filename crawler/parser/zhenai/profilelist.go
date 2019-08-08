package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strings"
)

const profileListRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`
const profileNextRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z/]+)">下一页</a>`

//城市个人信息列表解析
func ParseProfileList(body []byte) engine.ParseResult {
	body2 := string(body)
	compile := regexp.MustCompile(profileListRe)
	profileSubMatch := compile.FindAllStringSubmatch(body2, -1)
	//fmt.Println("profileSubMatch",profileSubMatch)
	parseResult := engine.ParseResult{}
	for _, item := range profileSubMatch {
		//addDataRequestQ(item[2], item[1], &parseResult)
		var profile model.Profile
		var url = item[1]
		profile.Name = item[2]
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: url,
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(url, bytes, profile)
			},
		})
		parseResult.Items = append(parseResult.Items, /* profile.Name*/ engine.Item{
			Url:     item[1],
			Id:      ExtractId(item[1]),
			Payload: profile.Name,
		})
	}
	//提取下一页数据
	nextPageSubMatch := regexp.MustCompile(profileNextRe).FindAllStringSubmatch(body2, -1)
	//fmt.Println("nextPageSubMatch",nextPageSubMatch)
	for _, v := range nextPageSubMatch {
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{Url: string(v[1]), ParserFunc: ParseProfileList})
		//parseResult.Items = append(parseResult.Items, string(v[2]))
	}
	return parseResult
}

func ExtractId(url string) string {
	var splits = strings.Split(url, "/")
	return splits[len(splits)-1]
}

/*func addDataRequestQ(name string, url string, parseResult *engine.ParseResult) {
	var profile model.Profile
	profile.Name = name
	parseResult.Requests = append(parseResult.Requests, engine.Request{
		Url: url,
		ParserFunc: func(bytes []byte) engine.ParseResult {
			return ParseProfile(bytes, profile)
		},
	})
	parseResult.Items = append(parseResult.Items,
		engine.Item{}
	profile.Name)
}*/
