package main

import (
	"crawler/web/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"web/view/template.html"))
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		panic(e)
	}

}
