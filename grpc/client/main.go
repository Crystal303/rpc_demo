package main

import (
	"context"
	"fmt"
	"io"
	"log"
	grpc2 "rpc_demo/grpc"
	"time"

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

	// 流处理
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 发送和接受两个操作
	go func() {
		for {
			if err := stream.Send(&grpc2.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

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
