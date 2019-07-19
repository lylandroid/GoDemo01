package main

import (
	"fmt"
	"time"
	"math/rand"
)

func worker(id int, c chan int) {
	func() {
		for v := range c {
			time.Sleep(2 * time.Second)
			fmt.Println(id, v)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Int31n(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	var values []int
	//10秒后会发送一个chan
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {

		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			//fmt.Println("c1", n)
			//w <- n
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
			//w <- n
			//fmt.Println("c2", n)
			//default:
			//	fmt.Println("No Value Received")
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("value len = ", len(values))
		case <-tm:
			fmt.Println("end")
			return
		}
	}

}
