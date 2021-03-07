package main

import (
	"github.com/mamachengcheng/12306/common"
	"github.com/mamachengcheng/12306/srv/payment/conf"
	"github.com/mamachengcheng/12306/srv/payment/handler"
	payment "github.com/mamachengcheng/12306/srv/payment/proto/payment"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
)

func main() {
	// Get Alipay Config
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")

	if err != nil {
		log.Fatalf("%v", err)
	}

	alipayConf := conf.GetAlipayFromConsul(consulConfig, "alipay")


	// Create service
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	srv := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	// Register handler
	payment.RegisterPaymentHandler(srv.Server(), &handler.Payment{AlipayConfig: alipayConf})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalf("Open MySQL database: %v", err)
	}
}
