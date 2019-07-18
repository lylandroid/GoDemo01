package main

import (
	"time"
	"fmt"
)

type Worker struct {
	in   chan int
	done chan bool
}

func worker2(id int, worker Worker) {
	func() {
		for v := range worker.in {
			fmt.Println(id, v)
			worker.done <- true
		}
	}()
}

//外部只能发数据：chan<- int，
// 外部只能收数据：<-chan int
func createWorker2(id int) Worker {
	w := Worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go worker2(id, w)
	return w
}

//通道demo
func chanDemo2() {
	var channels [10] Worker
	for i := 0; i < 10; i++ {
		channels[i] = createWorker2(i)
	}
	for i := 0; i < 10; i++ {
		channels[i].in <- 'a' + i
		<-channels[i].done
	}
	for i := 0; i < 10; i++ {
		channels[i].in <- 'A' + i
		<-channels[i].done
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
	chanDemo2()
	//bufferChannel()
	//channelClose()
}
