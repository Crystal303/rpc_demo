package impl

import (
	"context"
	grpc2 "rpc_demo/grpc"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
)

type PubSubService struct {
	grpc2.UnimplementedPubSubServiceServer
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p PubSubService) Subscribe(s *grpc2.String, server grpc2.PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			return strings.HasPrefix(key, s.GetValue())
		}
		return false
	})

	for v := range ch {
		if err := server.Send(&grpc2.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func (p PubSubService) Publish(ctx context.Context, s *grpc2.String) (*grpc2.String, error) {
	p.pub.Publish(s.GetValue())
	return &grpc2.String{}, nil
}
