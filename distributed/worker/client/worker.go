package client

import (
	"crawler/distributed/config"
	"crawler/distributed/worker"
	"crawler/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Process {
	return func(request engine.Request) engine.ParserResult {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParserResult
		c := <-clientChan
		e := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if e != nil {
			return engine.ParserResult{}
		}
		return worker.DeserializeResult(sResult)
	}
}
