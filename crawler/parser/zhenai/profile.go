package parser

import (
	"../../engine"
	"../../model"
	"regexp"
	"strconv"
)

const rootNodeProfileRe = `<div class="m-btn purple" [^>]*>([^<]+)</div>`

var compile = regexp.MustCompile(rootNodeProfileRe)
var compileNum = regexp.MustCompile("[0-9]+")

func ParseProfile(url string, body []byte, profile model.Profile) engine.ParseResult {
	submatch := compile.FindAllStringSubmatch(string(body), -1)
	index := 0
	//fmt.Println("submatch: ", len(submatch), submatch)
	profile.Marriage = string(submatch[index][1])
	index++
	profile.Age, _ = strconv.Atoi(compileNum.FindString(string(submatch[index][1])))
	index++
	profile.Xinzuo = string(submatch[index][1])
	index++
	profile.Height, _ = strconv.Atoi(compileNum.FindString(string(submatch[index][1])))
	if len(submatch) == 9 {
		index++
		profile.Weight, _ = strconv.Atoi(compileNum.FindString(string(submatch[index][1])))
	}
	index++
	profile.HuKou = string(submatch[index][1])
	index++
	profile.Income = string(submatch[index][1])
	index++
	profile.Occupation = string(submatch[index][1])
	index++
	if len(submatch) < index {
		profile.Education = string(submatch[index][1])
	}

	parseResult := engine.ParseResult{}
	parseResult.Items = append(parseResult.Items,
		engine.Item{
			Url:     url,
			Type:    "zhenai",
			Id:      ExtractId(url),
			Payload: profile,
		})
	return parseResult

}
