package engine

import (
	"../fetcher"
	"fmt"
)

func Run(requests ...Request) {
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		fmt.Println("request url：", r.Url)
		if err != nil {
			fmt.Printf("Fetcher: error url %s   %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			fmt.Printf("Got item %v", item)
			fmt.Println()
		}
	}
}
