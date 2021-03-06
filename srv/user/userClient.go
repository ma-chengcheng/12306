package main

import (
	"context"
	goMicroServiceUser "github.com/mamachengcheng/12306/srv/user/proto/user"
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
		micro.Name("go.micro.service.user.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	userService := goMicroServiceUser.NewUserService("go.micro.service.user", srv.Client())

	reply, err := userService.Login(context.TODO(), &goMicroServiceUser.LoginRequest{
		Username: "test",
		Password: "test",
	})

	//reply, err :=  userService.Register(context.TODO(), &goMicroServiceUser.RegisterRequest{
	//	Username:    "test",
	//	Password:    "test",
	//	Email:       "123@qq.com",
	//	MobilePhone: "15608652021",
	//	Name:        "xu",
	//	Certificate: "430426199811145139",
	//})

	//	BookTickets(context.TODO(), &goMicroServiceUser.BookTicketsRequest{
	//	OrderID:     0,
	//	ScheduleID:  0,
	//	SeatType:    0,
	//	PassengerID: nil,
	//})

	if err != nil {
		log.Fatalf("%v", err)
		//log.Println("1")
	}

	log.Printf("%v", reply.IsSuccess)
	log.Printf("%v", reply.Msg)
	//log.Println("2")
}
