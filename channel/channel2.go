package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in   chan int
	done func()
}

func worker2(id int, worker Worker) {
	func() {
		for v := range worker.in {
			fmt.Println(id, string(v))
			worker.done()
		}
	}()
}

//外部只能发数据：chan<- int，
// 外部只能收数据：<-chan int
func createWorker2(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go worker2(id, w)
	return w
}

//通道demo
func chanDemo2() {
	var wg sync.WaitGroup
	var workers [10]Worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}
	for i, w := range workers {
		w.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo2()
	//bufferChannel()
	//channelClose()
}
