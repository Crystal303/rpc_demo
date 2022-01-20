package main

import (
	"log"
	"net"
	grpc2 "rpc_demo/grpc"
	"rpc_demo/grpc/impl"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	grpc2.RegisterHelloServiceServer(grpcServer, impl.NewHelloServiceServer())
	grpc2.RegisterPubSubServiceServer(grpcServer, impl.NewPubSubService())

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
