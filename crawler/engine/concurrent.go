package engine

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan() chan Request
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//fmt.Printf("Got item: %v\n", item)
			go func(item interface{}) {
				e.ItemChan <- item
			}(item)
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			if parseResult, err := worker(request); err != nil {
				continue
			} else {
				out <- parseResult
			}
		}
	}()
}
