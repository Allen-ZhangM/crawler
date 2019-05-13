package parser

import (
	"crawler/engine"
	"regexp"
)

var regexCity = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(bytes []byte) engine.ParserResult {
	all := regexCity.FindAllSubmatch(bytes, -1)

	result := engine.ParserResult{}
	i := 0
	for _, m := range all {
		i++
		if i > 500 {
			break
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
