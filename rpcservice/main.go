package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc_demo/rpcservice/impl"
)

func main() {
	_ = RegisterHelloService(new(impl.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp err: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept err: ", err)
		}
		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
