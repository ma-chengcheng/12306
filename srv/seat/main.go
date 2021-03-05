package main

import (
	"seat/handler"
	pb "seat/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("seat"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterSeatHandler(srv.Server(), new(handler.Seat))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
