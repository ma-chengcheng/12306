package main

import (
	"github.com/mamachengcheng/12306/srv/user/common"
	"github.com/mamachengcheng/12306/srv/user/domain/respository"
	s "github.com/mamachengcheng/12306/srv/user/domain/service"
	"github.com/mamachengcheng/12306/srv/user/handler"
	user "github.com/mamachengcheng/12306/srv/user/proto"
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
		micro.Name("userAPI"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	userDataService := s.NewUserDataService(respository.NewUserRepository(db))

	// Register handler
	user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalf("Open MySQL database: %v", err)
	}
}
