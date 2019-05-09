package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var regexAge = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

//var regexProfile = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^|]+) | ([0-9]+)岁 | ([^|]+) | ([^|]+) | ([^|]+) | ([^<]+)</div>`)
var regexProfile = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^<]+)</div>`)
var regexGender = regexp.MustCompile(`"genderString":"([^"]+)"`)

func ParseProfile(bytes []byte, photoUrl string, name string) engine.ParserResult {
	allProfile := regexProfile.FindAllSubmatch(bytes, -1)
	allGender := regexGender.FindAllSubmatch(bytes, -1)

	var result engine.ParserResult
	for i, m := range allGender {
		//fmt.Println(string(m[1]))
		//result.Requests = append(result.Requests,engine.Request{
		//	Url:string(m[1]),
		//})
		//age, _ := strconv.Atoi(string(m[2]))
		result.Item = append(result.Item, model.Profile{
			Name:   name,
			Gender: string(m[1]),
			//Age:    0,
			Education: string(allProfile[i][1]),
			//Marriage:string(m[4]),
			//Height:string(m[5]),
			//Income:string(m[6]),
		})
	}

	return result
}
