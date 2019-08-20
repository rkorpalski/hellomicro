package service

import (
	"context"
	proto "github.com/rkorpalski/hellomicro/proto"
)

type HelloService struct {}

func (s *HelloService) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func NewHelloService() *HelloService {
	return &HelloService{}
}
