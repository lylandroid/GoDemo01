package engine

import (
	"../fetcher"
	"fmt"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(requests ...Request) {
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		if parseResult, err := worker(r); err != nil {
			continue
		} else {
			requests = append(requests, parseResult.Requests...)
			for _, item := range parseResult.Items {
				fmt.Printf("Got item %v", item)
				fmt.Println()
			}
		}
	}
}

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	fmt.Println("request url：", r.Url)
	if err != nil {
		fmt.Printf("Fetcher: error url %s   %v", r.Url, err)
		fmt.Println()
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
