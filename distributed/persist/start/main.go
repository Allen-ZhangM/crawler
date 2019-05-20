package main

import (
	"crawler/distributed/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenaiwang/parser"
)

func main() {
	runConcurrent()
}

func runConcurrent() {
	client := client.ItemSaver(":1234")

	e := engine.ConcurrentEngine{
		Scheduler:   &shceduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    client,
	}
	e.RunConcurrentRequest(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/yingkou",
		ParserFunc: parser.ParseCityList,
	})
}
