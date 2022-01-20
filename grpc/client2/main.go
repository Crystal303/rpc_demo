package main

import (
	"context"
	"fmt"
	"io"
	"log"
	grpc2 "rpc_demo/grpc"

	"google.golang.org/grpc"
)

// main 订阅消息客户端
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc2.NewPubSubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &grpc2.String{Value: "golang"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
