package test

import (
	"context"
	goMicroServiceSeat "github.com/mamachengcheng/12306/srv/seat/proto/seat"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := micro.NewService(
		micro.Name("go.micro.service.payment.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	ticketService := goMicroServiceSeat.NewSeatService("go.micro.service.seat", srv.Client())

	cuntRemainingSeatsReply, err := ticketService.CountRemainingSeats(context.TODO(), &goMicroServiceSeat.CountRemainingSeatsRequest{
		SeatType:       0,
		TrainID:        0,
		ScheduleStatus: 0,
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", cuntRemainingSeatsReply.Number)
	}
}
