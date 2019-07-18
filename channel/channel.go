package main

import (
	"time"
	"fmt"
)

func worker(id int, c chan int) {
	func() {
		/*for {
			if n, ok := <-c; ok {
				fmt.Println(id, n)
			} else {
				break
			}
		}*/
		for v := range c {
			fmt.Println(id, v)
		}
	}()
}

//外部只能发数据：chan<- int，
// 外部只能收数据：<-chan int
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond * 10)
}

func chanDemo() {
	var channels [10] chan<- int
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

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond * 10)

}

func main() {
	//chanDemo()
	//bufferChannel()
	channelClose()
}
