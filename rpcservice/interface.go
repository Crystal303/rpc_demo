package main

import (
	"net/rpc"
	"rpc_demo/constant"
)

type HelloInterface = interface {
	Hello(string, *string) error
}

func RegisterHelloService(svc HelloInterface) error {
	return rpc.RegisterName(constant.HelloServiceName, svc)
}
