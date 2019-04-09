package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(bytes []byte) engine.ParserResult {
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllSubmatch(bytes, -1)

	result := engine.ParserResult{}
	for _, m := range all {
		result.Item = append(result.Item, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParserFunc,
		})
	}
	return result
}
