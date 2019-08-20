package main

import (
	"context"
	"github.com/micro/go-micro"
	proto "github.com/rkorpalski/hellomicro/proto"
	"log"
)

var clientName = "client-helloworld"
var serviceName = "service-helloworld"

func main() {

	service := micro.NewService(micro.Name(clientName))
	service.Init()

	hello := proto.NewGreeterClient(serviceName, service.Client())

	rsp, err := hello.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		log.Println(err)
	}

	log.Println(rsp.Greeting)
}


