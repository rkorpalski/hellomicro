package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/rkorpalski/hellomicro/proto"
	"log"
	"time"
	"github.com/pborman/uuid"
)

var clientName = "client-helloworld"

func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {

		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Printf("publishing %+v\n", ev)


		if err := p.Publish(context.Background(), ev); err != nil {
			log.Printf("error publishing: %v", err)
		}
	}
}


func main() {

	service := micro.NewService(micro.Name(clientName))
	service.Init()

	pub1 := micro.NewPublisher("topic.service", service.Client())

	go sendEv("topic.service", pub1)


	select {}
}