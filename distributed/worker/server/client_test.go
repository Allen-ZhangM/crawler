package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcSupport"
	"crawler/distributed/worker"
	"crawler/zhenaiwang/parser"
	"fmt"
	"testing"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcSupport.ServerRPC(host, worker.CrawlService{})

	c, _ := rpcSupport.NewClient(host)

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1737015172",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: parser.ProfileParser{
				UserName: "UserName",
				PhotoUrl: "PhotoUrl",
			},
		},
	}
	var result = worker.ParserResult{}

	e := c.Call(config.CrawlServiceRpc, req, &result)

	if e != nil {
		t.Error(e)
	} else {
		fmt.Println(result)
	}

}
