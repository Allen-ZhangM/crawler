package parser

import (
	"crawler/engine"
	"regexp"
)

var regexName = regexp.MustCompile(`<div class="photo"><a href="(http://album.zhenai.com/u/[0-9a-z]+)" target="_blank"><img src="(https://photo.zastatic.com/images/photo/[^"]*)" alt="([^"]+)"></a></div>`)
var regexNextPage = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+/[0-9]+)">下一页</a>`)

func ParseCity(bytes []byte, url string) engine.ParserResult {
	all := regexName.FindAllSubmatch(bytes, -1)
	allRegexNextPage := regexNextPage.FindAllSubmatch(bytes, -1)
	result := engine.ParserResult{}
	for _, m := range all {
		name := string(m[3])
		photoUrl := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: NewProfileParser(name, photoUrl),
		})
	}

	for _, m := range allRegexNextPage {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}
	return result
}
