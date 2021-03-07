package main

import (
	"order/handler"
	pb "order/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("order"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterOrderHandler(srv.Server(), new(handler.Order))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
