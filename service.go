package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/Nanixel/FirstDownMicro/statsservice/proto"
	"github.com/micro/go-micro/broker"
)

const topic = "game.runplay"

type statsservice struct {
	PubSub broker.Broker
}

func (s *statsservice) Run() {

	if connerr := s.PubSub.Connect(); connerr != nil {
		log.Fatal(connerr)
	}

	_, e := s.PubSub.Subscribe(topic, func(p broker.Publication) error {
		var m map[string]string
		if err := json.Unmarshal(p.Message().Body, &m); err != nil {
			fmt.Println("Error occured when unmarshaling.")
			return err
		}
		fmt.Println(m["name"])
		return nil
	})

	if e != nil {
		fmt.Println(e)
	}
}

func (s *statsservice) GetPlayerBaseStats(ctx context.Context, req *pb.GetPlayerStatsRequest, res *pb.BaseStatsResponse) error {

	fmt.Println("GetPlayerBaseStatesCalled")

	return nil
}

// This will be called on the subscribe
func (s *statsservice) UpdatePlayerStats() {

}
