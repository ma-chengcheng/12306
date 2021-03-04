package main

import (
	"context"
	goMicroServiceTicket "github.com/mamachengcheng/12306/srv/ticket/proto/ticket"
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
		micro.Name("go.micro.service.ticket.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	ticketService := goMicroServiceTicket.NewTicketService("go.micro.service.ticket", srv.Client())

	reply, err :=  ticketService.BookTickets(context.TODO(), &goMicroServiceTicket.BookTicketsRequest{
		OrderID:     0,
		ScheduleID:  0,
		SeatType:    0,
		PassengerID: nil,
	})

	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%v", reply.IsSuccess)
}
