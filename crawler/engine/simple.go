package engine

import (
	"fmt"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(requests ...Request) {
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		if parseResult, err := Worker(r); err != nil {
			continue
		} else {
			requests = append(requests, parseResult.Requests...)
			for _, item := range parseResult.Items {
				fmt.Printf("Got item %v\n", item)
			}
		}
	}
}
