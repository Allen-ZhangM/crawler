package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strings"
)

var regexAge = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

//var regexProfile = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^|]+) | ([0-9]+)岁 | ([^|]+) | ([^|]+) | ([^|]+) | ([^<]+)</div>`)
var regexProfile = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^<]+)</div>`)
var regexGender = regexp.MustCompile(`"genderString":"([^"]+)"`)

func ParseProfile(bytes []byte, photoUrl string, name string, url string) engine.ParserResult {
	allProfile := regexProfile.FindAllSubmatch(bytes, -1)
	allGender := regexGender.FindAllSubmatch(bytes, -1)

	urlSplits := strings.Split(url, "/")

	var result engine.ParserResult
	for i, m := range allGender {
		result.Item = append(result.Item, engine.Item{
			Url:  url,
			Type: Type,
			Id:   urlSplits[len(urlSplits)-1],
			Payload: model.Profile{
				Name:      name,
				PhotoUrl:  photoUrl,
				Gender:    string(m[1]),
				Education: string(allProfile[i][1]),
			},
		})
	}

	return result
}
