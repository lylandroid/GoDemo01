package engine

import (
	"../fetcher"
	"fmt"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	fmt.Println("request url：", r.Url)
	if err != nil {
		fmt.Printf("Fetcher: error url %s   %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
