package main

import (
	"fmt"

	pb "github.com/Nanixel/FirstDownMicro/statsservice/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
)

// type service struct{}

// const topic = "game.runplay"

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.playerservice"),
		micro.Version("latest"),
	)
	service.Init()

	pubsub := service.Server().Options().Broker
	srv := &statsservice{PubSub: pubsub}

	pb.RegisterStatsServiceHandler(service.Server(), srv)

	// if connerr := pubsub.Connect(); connerr != nil {
	// 	log.Fatal(connerr)
	// }

	// _, e := pubsub.Subscribe(topic, func(p broker.Publication) error {
	// 	var m map[string]string
	// 	if err := json.Unmarshal(p.Message().Body, &m); err != nil {
	// 		fmt.Println("Error occured when unmarshaling.")
	// 		return err
	// 	}
	// 	fmt.Println(m["name"])
	// 	return nil
	// })

	// if e != nil {
	// 	fmt.Println(e)
	// }

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
