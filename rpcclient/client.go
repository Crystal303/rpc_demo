package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc_demo/constant"
)

type HelloServiceClient struct {
	*rpc.Client
}

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	//return &HelloServiceClient{rpc.NewClient(conn)}, nil
	return &HelloServiceClient{rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}, nil
}

func (c *HelloServiceClient) Hello(request string, reply *string) error {
	if c.Client == nil {
		return nil
	}
	return c.Client.Call(constant.HelloServiceName+".Hello", request, reply)
}
