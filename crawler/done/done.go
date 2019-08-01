package done

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{}
}
