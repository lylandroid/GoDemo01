package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine  i=%d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)

}
