package impl

import (
	"context"
	"io"
	"rpc_demo/grpc"
)

var _ grpc.HelloServiceServer = (*HelloServiceImpl)(nil)

func NewHelloServiceServer() grpc.HelloServiceServer {
	return new(HelloServiceImpl)
}

type HelloServiceImpl struct {
	grpc.UnimplementedHelloServiceServer
}

func (h *HelloServiceImpl) Hello(ctx context.Context, s *grpc.String) (*grpc.String, error) {
	reply := &grpc.String{Value: "hello:" + s.GetValue()}
	return reply, nil
}

func (h *HelloServiceImpl) Channel(server grpc.HelloService_ChannelServer) error {
	for {
		args, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &grpc.String{Value: "hello:" + args.GetValue()}
		if err := server.Send(reply); err != nil {
			return err
		}
	}
}
