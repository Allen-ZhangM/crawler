package worker

import "crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(request Request, result *ParserResult) error {
	engineReq, e := DeserializeRequest(request)
	if e != nil {
		return e
	}

	engineResult := engine.Worker(engineReq)

	*result = SerializeResult(engineResult)

	return nil
}
