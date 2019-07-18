package main

import (
	"time"
	"fmt"
)

func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Println(id, n)
	}
}

func chanDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
