package engine

import (
	"crawler/fetcher"
	"log"
)

func RunFetch(seeds ...Request) {
	var queue []Request
	for _, seed := range seeds {
		queue = append(queue, seed)
	}

	for len(queue) > 0 {
		request := queue[0]
		queue = queue[1:]

		log.Printf("Fetching url: %s", request.Url)
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", request.Url, err)
			continue
		}
		parseResult := request.ParserFunc(body)
		queue = append(queue, parseResult.Requests...)

		for _, item := range parseResult.Item {
			log.Printf("%s", item)
		}

	}

}

func RunRequest(seeds ...Request) {
	var queue []Request
	for _, seed := range seeds {
		queue = append(queue, seed)
	}
	i := 0
	for len(queue) > 0 {
		request := queue[0]
		queue = queue[1:]

		log.Printf("Request url: %s", request.Url)

		body, err := fetcher.Request(request.Url)

		if err != nil {
			log.Printf("Request: error Request url %s: %v", request.Url, err)
			continue
		}

		parseResult := request.ParserFunc(body)
		queue = append(queue, parseResult.Requests...)

		for _, item := range parseResult.Item {
			log.Printf("get item %d: %s", i, item)
			i++
		}

	}

}
