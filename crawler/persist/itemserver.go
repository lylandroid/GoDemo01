package persist

import "fmt"

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount ++
			fmt.Printf("Item Saves: %d %v\n", itemCount, item)
		}
	}()
	return out
}
