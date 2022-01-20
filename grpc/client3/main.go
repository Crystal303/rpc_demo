package main

import (
	"context"
	"log"
	grpc2 "rpc_demo/grpc"

	"google.golang.org/grpc"
)

// main 负责发布消息客户端，发布时间点要在订阅之后
func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc2.NewPubSubServiceClient(conn)
	_, err = client.Publish(context.Background(), &grpc2.String{Value: "golang: pubSub"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &grpc2.String{Value: "docker: pubSub"})
	if err != nil {
		log.Fatal(err)
	}
}
