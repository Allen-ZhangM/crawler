package controller

import (
	"context"
	"crawler/engine"
	"crawler/web/model"
	"crawler/web/view"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchHandler struct {
	view   view.SearchResult
	client *elastic.Client
}

func (s SearchHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, e := strconv.Atoi(req.FormValue("from"))
	if e != nil {
		from = 0
	}

	var page model.SearchResult
	page, e = s.getSearchResult(q, from)
	if e != nil {
		http.Error(res, e.Error(), http.StatusBadRequest)
	}
	fmt.Println(page)
	e = s.view.Render(res, page)
	if e != nil {
		http.Error(res, e.Error(), http.StatusBadRequest)
	}
}

func (s *SearchHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	var resp *elastic.SearchResult
	var err error
	if q == "" {
		resp, err = s.client.Search("dating_profile").From(from).Do(context.Background())
	} else {
		resp, err = s.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())
	}
	if err != nil {
		return result, err
	}
	fmt.Println(resp.TotalHits())
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func CreateSearchResultHandler(template string) SearchHandler {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.245.137:9200"),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}
