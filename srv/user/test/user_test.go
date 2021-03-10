package test

import (
	"context"
	goMicroServiceUser "github.com/mamachengcheng/12306/srv/user/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
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

	registerReply, err := userService.Register(context.TODO(), &goMicroServiceUser.RegisterRequest{
		Username:    "test2",
		Password:    "test2",
		Email:       "456@qq.com",
		MobilePhone: "13397349018",
		Name:        "xiao",
		Certificate: "430426196712015116",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", registerReply.Msg)
		log.Printf("%v", registerReply.IsSuccess)
	}
}

func TestLogin(t *testing.T) {
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

	loginReply, err := userService.Login(context.TODO(), &goMicroServiceUser.LoginRequest{
		Username: "test2",
		Password: "test2",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", loginReply.Msg)
		log.Printf("%v", loginReply.IsSuccess)
	}
}

func TestQueryUserInformation(t *testing.T) {
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

	queryUserInformationReply, err := userService.QueryUserInformation(context.TODO(), &goMicroServiceUser.QueryUserInformationRequest{
		Username: "test2",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", queryUserInformationReply)
	}
}

func TestUpdatePassword(t *testing.T) {
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

	updatePasswordReply, err := userService.UpdatePassword(context.TODO(), &goMicroServiceUser.UpdatePasswordRequest{
		Username: "test2",
		Password: "test3",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", updatePasswordReply.IsSuccess)
		log.Printf("%v", updatePasswordReply.Msg)
	}
}

func TestAddRegularPassenger(t *testing.T) {
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

	addRegularPassengerReply, err := userService.AddRegularPassenger(context.TODO(), &goMicroServiceUser.AddRegularPassengerRequest{
		Username:    "test2",
		MobilePhone: "17671434306",
		Name:        "xiu",
		Certificate: "430426197010125120",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", addRegularPassengerReply.IsSuccess)
		log.Printf("%v", addRegularPassengerReply.Msg)
	}
}

func TestQueryRegularPassengers(t *testing.T) {
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

	queryRegularPassengersReply, err := userService.QueryRegularPassengers(context.TODO(), &goMicroServiceUser.QueryRegularPassengersRequest{
		Username: "test2",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", queryRegularPassengersReply)
	}
}

func TestUpdateRegularPassenger(t *testing.T) {
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

	updateRegularPassengerReply, err := userService.UpdateRegularPassenger(context.TODO(), &goMicroServiceUser.UpdateRegularPassengerRequest{
		Username:      "test2",
		PassengerID:   "11",
		MobilePhone:   "17671434307",
		PassengerType: "1",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", updateRegularPassengerReply.IsSuccess)
		log.Printf("%v", updateRegularPassengerReply.Msg)
	}
}

func TestDeleteRegularPassenger(t *testing.T) {
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

	deleteRegularPassengerReply, err := userService.DeleteRegularPassenger(context.TODO(), &goMicroServiceUser.DeleteRegularPassengerRequest{
		PassengerID: "11",
	})
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("%v", deleteRegularPassengerReply.IsSuccess)
		log.Printf("%v", deleteRegularPassengerReply.Msg)
	}
}
