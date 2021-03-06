package main

import (
	"context"
	goMicroServiceSeat "github.com/mamachengcheng/12306/srv/seat/proto/seat"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

func main() {

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := micro.NewService(
		micro.Name("go.micro.service.seat.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	ticketService := goMicroServiceSeat.NewSeatService("go.micro.service.seat", srv.Client())

	countRemainingSeatsReply, _ := ticketService.CountRemainingSeats(context.TODO(), &goMicroServiceSeat.CountRemainingSeatsRequest{
		SeatType:       0,
		TrainID:        1,
		ScheduleStatus: 3,
	})

	log.Printf("%v", countRemainingSeatsReply.Number)

	_, err := ticketService.GetSeats(context.TODO(), &goMicroServiceSeat.GetSeatsRequest{
		SeatType:       1,
		ScheduleStatus: 3,
		Number:         1,
		TrainID:        1,
	})

	log.Printf("%v", err)

}
