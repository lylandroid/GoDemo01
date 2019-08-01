package engine

import "fmt"

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v\n", item)
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			if parseResult, err := worker(request); err != nil {
				continue
			} else {
				out <- parseResult
			}
		}
	}()
}
