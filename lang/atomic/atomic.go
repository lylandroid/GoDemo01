package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	return int(a.value)
}

func main() {
	var a atomicInt
	for i := 0; i < 300; i++ {
		a.increment()
		go func() {
			a.increment()
		}()
	}
	fmt.Println(a.get())
	time.Sleep(time.Second * 10)
	fmt.Println(a.get())
}
