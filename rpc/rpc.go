package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRPC(host string, service interface{}) {
	rpc.Register(service)

	l, e := net.Listen("tcp", host)
	if e != nil {
		panic(e)
	}

	for {
		c, e := l.Accept()
		if e != nil {
			log.Println("Accept err:", e)
			continue
		}
		go jsonrpc.ServeConn(c)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	c, e := net.Dial("tcp", host)
	if e != nil {
		return nil, e
	}
	return jsonrpc.NewClient(c), nil
}
