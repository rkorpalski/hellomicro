package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	proto "github.com/rkorpalski/hellomicro/proto"
	"github.com/rkorpalski/hellomicro/service"
	hystrixplugin "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"github.com/afex/hystrix-go/hystrix"
	"log"
)

var serviceName = "service-helloworld"

func main() {

	// time for waiting command response
	hystrix.DefaultTimeout = 1200
	// how long to open circuit breaker again
	hystrix.DefaultSleepWindow = 2000
	// percent of bad response
	hystrix.DefaultErrorPercentThreshold = 10
	// how much request can be accessed in persecond
	hystrix.DefaultMaxConcurrent = 2
	//how much requests to enable circuit breaker
	hystrix.DefaultVolumeThreshold = 1


	server := micro.NewService(
		micro.Name(serviceName),
		micro.WrapClient(hystrixplugin.NewClientWrapper()),
		micro.Server(
			server.NewServer(
				server.Name(serviceName),
				server.Address(":8084"),
			),
		),
	)

	proto.RegisterGreeterHandler(server.Server(), service.NewHelloService())

	server.Init()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
