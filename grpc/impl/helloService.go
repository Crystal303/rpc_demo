package impl

import (
	"context"
	"io"
	"rpc_demo/grpc"

	"google.golang.org/grpc/metadata"
)

var _ grpc.HelloService_ChannelServer = (*HelloServiceImpl)(nil)

var _ io.Reader = (*HelloServiceImpl)(nil)

func NewHelloService() io.Reader {
	return new(HelloServiceImpl)
}

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
	panic("implement me")
}

func (h *HelloServiceImpl) Send(s *grpc.String) error {
	panic("implement me")
}

func (h *HelloServiceImpl) Recv() (*grpc.String, error) {
	panic("implement me")
}

func (h *HelloServiceImpl) SetHeader(md metadata.MD) error {
	panic("implement me")
}

func (h *HelloServiceImpl) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (h *HelloServiceImpl) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (h *HelloServiceImpl) Context() context.Context {
	panic("implement me")
}

func (h *HelloServiceImpl) SendMsg(m interface{}) error {
	panic("implement me")
}

func (h *HelloServiceImpl) RecvMsg(m interface{}) error {
	panic("implement me")
}

func (h *HelloServiceImpl) Read(p []byte) (n int, err error) {
	panic("implement me")
}
