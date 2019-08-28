package rpcsupport

import (
	"../../crawlers_distributed/worker"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, openServerApi interface{}) error {
	rpc.Register(openServerApi)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	} else {
		log.Printf("Listening  on %s", host)
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

func (app *AppRpcClient) NewRpcClient(host string) error {
	client, err := NewRpcClient(host)
	if err != nil {
		return err
	}
	app.Client = client
	return nil
}

func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	var client *rpc.Client
	client = jsonrpc.NewClient(conn)
	return client, nil
}

func (client AppRpcClient) CallFun(funcName string, args interface{}) (string, error) {
	var result string
	return result, client.Client.Call(funcName, args, &result)
}

func (client AppRpcClient) CallFun2(funcName string, args interface{}) (worker.SerializeParseResult, error) {
	var result worker.SerializeParseResult
	return result, client.Client.Call(funcName, args, &result)
}
