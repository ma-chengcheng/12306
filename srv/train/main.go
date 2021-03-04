package main

import (
	"github.com/mamachengcheng/12306/services/train/common"
	"github.com/mamachengcheng/12306/services/train/domain/respository"
	s "github.com/mamachengcheng/12306/srv/train/domain/service"
	"github.com/mamachengcheng/12306/services/train/handler"
	train "github.com/mamachengcheng/12306/srv/train/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

func main() {

	db, err := common.GetMySqlDB()
	if err != nil {
		log.Fatalf("Open MySQL database: %v", err)
	}

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("train"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	trainDataService := s.NewTrainDataService(respository.NewTrainRepository(db))

	// Register handler
	train.RegisterTrainHandler(srv.Server(), &handler.Train{TrainDataService: trainDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalf("Open MySQL database: %v", err)
	}
}
