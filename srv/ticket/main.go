package main

import (
	"github.com/mamachengcheng/12306/services/ticket/domain/respository"
	s "github.com/mamachengcheng/12306/srv/ticket/domain/service"
	"github.com/mamachengcheng/12306/services/ticket/handler"
	ticket "github.com/mamachengcheng/12306/srv/ticket/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

func main() {

	//if err != nil {
	//	log.Fatalf("Open MySQL database: %v", err)
	//}

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("userAPI"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	userDataService := s.NewTicketDataService(respository.NewTicketRepository(db))

	// Register handler

	ticket.RegisterTicketHandler(srv.Server(), &handler.Ticket{TicketDataService: userDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalf("Open MySQL database: %v", err)
	}
}
