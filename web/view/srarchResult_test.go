package view

import (
	"crawler/engine"
	m "crawler/model"
	"crawler/web/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 124
	page.Start = 10
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1358992404",
		Type: "zhenaiwang",
		Id:   "1358992404",
		Payload: m.Profile{
			Name:           "白雪王子",
			Gender:         "男士",
			Age:            0,
			Height:         "",
			Weight:         "",
			Income:         "",
			Marriage:       "",
			Education:      "阿坝 | 45岁 | 高中及以下 | 未婚 | 175cm | 3000元以下",
			Occupation:     "",
			CensusRegister: "",
			Constellation:  "",
			Horse:          "",
			Car:            "",
			PhotoUrl:       "https://photo.zastatic.com/images/photo/339749/1358992404/6409106887960068.jpg?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north",
		},
	}
	page.Items = append(page.Items, item)
	page.Items = append(page.Items, item)
	page.Items = append(page.Items, item)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
