package engine

import "crawler/distributed/config"

type Request struct {
	Url    string
	Parser Parser
}

type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

type ParserResult struct {
	Requests []Request
	Item     []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParserFunc([]byte) ParserResult {
	return ParserResult{}
}

type NilParser struct{}

func (NilParser) Parse(contents []byte, url string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type ParserFunc func(contents []byte, url string) ParserResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, args
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
