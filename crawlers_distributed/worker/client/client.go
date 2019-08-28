package client

import (
	"../../../crawler/engine"
	"../../config"
	"../../worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	/*appRpcClient := rpcsupport.AppRpcClient{}
	if err := appRpcClient.NewRpcClient(config.AppAddress); err != nil {
		return nil, err
	}*/
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		client := <-clientChan
		var sParseResult worker.SerializeParseResult
		err := client.Call(config.CrawlersServiceSaverRPCApi, sReq, &sParseResult)
		//sParseResult, err := appRpcClient.CallFun2(config.CrawlersServiceSaverRPCApi, sReq)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeResult(sParseResult), nil
	}

}
