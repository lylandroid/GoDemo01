package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"../../crawlers_distributed/worker"
)

func ServeRpc(host string, openServerApi interface{}) error {
	rpc.Register(openServerApi)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

type AppRpcClient struct {
	Client *rpc.Client
}

func (client *AppRpcClient) NewRpcClient(host string) error {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}
	client.Client = jsonrpc.NewClient(conn)
	return nil
}

func (client AppRpcClient) CallFun(funcName string, args interface{}) (string, error) {
	var result string
	return result, client.Client.Call(funcName, args, &result)
}

func (client AppRpcClient) CallFun2(funcName string, args interface{}) (worker.SerializeParseResult, error) {
	var result worker.SerializeParseResult
	return result, client.Client.Call(funcName, args, &result)
}
