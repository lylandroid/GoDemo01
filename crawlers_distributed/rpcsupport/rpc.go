package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, server interface{}) error {
	rpc.Register(server)
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

func (client AppRpcClient) CallFun2(funcName string, args interface{}) (interface{}, error) {
	var result interface{}
	return result, client.Client.Call(funcName, args, &result)
}
