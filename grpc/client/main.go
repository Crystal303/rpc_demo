package main

import (
	"context"
	"fmt"
	"log"
	grpc2 "rpc_demo/grpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc2.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &grpc2.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
