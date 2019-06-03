package worker

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/zhenaiwang/parser"
	"github.com/goinggo/mapstructure"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParserResult {
	result := ParserResult{
		Items: r.Item,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, e := DeserializeParser(r.Parser)
	if e != nil {
		return engine.Request{}, e
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(r ParserResult) engine.ParserResult {
	result := engine.ParserResult{
		Item: r.Items,
	}
	for _, req := range r.Requests {
		r, e := DeserializeRequest(req)
		if e != nil {
			log.Println("DeserializeResult error:", e)
			continue
		}
		result.Requests = append(result.Requests, r)
	}
	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		pp := parser.ProfileParser{}
		e := mapstructure.Decode(p.Args, &pp)
		if e != nil {
			log.Println(e)
		}
		return parser.NewProfileParser(pp.UserName, pp.PhotoUrl), nil
	default:
		return nil, errors.New("unknown parser name:" + p.Name)
	}
}
