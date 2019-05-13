package engine

import (
	"crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) RunConcurrentRequest(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	for {
		result := <-out
		for _, item := range result.Item {
			e.ItemChan <- item
		}

		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
			//in <- req
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, r ReadyNotifier) {
	go func() {
		for {
			r.WorkerReady(in)
			result := <-in
			parserResult := worker(result)
			out <- parserResult
		}
	}()
}

func worker(request Request) ParserResult {
	log.Printf("Request url: %s", request.Url)

	body, err := fetcher.Request(request.Url)

	if err != nil {
		log.Printf("Request: error Request url %s: %v", request.Url, err)
		return ParserResult{}
	}

	return request.ParserFunc(body)
}
