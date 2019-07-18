package main

import (
	"time"
	"fmt"
)

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(id, (string(<-c)))
		}
	}()
	return c
}

func chanDemo() {
	var channels [10] chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond * 10)
	//Output
	//0 97
	//4 101
	//1 98
	//2 99
	//3 100
	//6 103
	//5 102
	//7 104
	//8 105
	//9 106
}

func main() {
	chanDemo()
}
