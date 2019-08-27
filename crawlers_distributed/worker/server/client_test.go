package main

import (
	"fmt"
	"testing"
	"../../rpcsupport"
	"../../worker"
	"time"
	"../../config"
)

func TestCrawlersService(t *testing.T) {
	const host = ":9002"
	go rpcsupport.ServeRpc(host, worker.CrawlersService{})
	time.Sleep(time.Second)

	appRpcClient := rpcsupport.AppRpcClient{}
	appRpcClient.NewRpcClient(host)
	result, err := appRpcClient.CallFun2(config.CrawlersServiceSaverRPCApi, worker.Request{
		Url: "http://album.zhenai.com/u/1511101827",
		Parser: worker.SerializedParser{
			FuncName: config.FuncParseProfile,
			Args:     "雪儿只聊天",
		},
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

}
