package worker

import (
	"../../crawler/engine"
)

type CrawlersService struct {
}

func (CrawlersService) Process(sReq Request, result *SerializeParseResult) error {
	req, err := DeserializeRequest(sReq)
	if err != nil {
		return err
	}
	parseResult, err := engine.Worker(req)
	if err != nil {
		return err
	}
	*result = SerializeResult(parseResult)
	return nil
}
