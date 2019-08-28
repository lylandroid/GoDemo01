package client

import (
	"../../rpcsupport"
	"../../config"
	"../../../crawler/engine"
	"../../worker"
)

func CreateProcessor() (engine.Processor, error) {
	appRpcClient := rpcsupport.AppRpcClient{}
	if err := appRpcClient.NewRpcClient(config.AppAddress); err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		sParseResult, err := appRpcClient.CallFun2(config.CrawlersServiceSaverRPCApi, sReq)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sParseResult), nil

	}, nil

}
