package view

import (
	"crawler/web/model"
	"html/template"
	"io"
)

type SearchResult struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResult {
	return SearchResult{
		template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResult) Render(wr io.Writer, data model.SearchResult) error {
	return s.template.Execute(wr, data)
}
