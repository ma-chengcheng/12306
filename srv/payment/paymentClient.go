package main

import (
	"context"
	goMicroServicePayment "github.com/mamachengcheng/12306/srv/payment/proto/payment"
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
		micro.Name("go.micro.service.payment.client"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	ticketService := goMicroServicePayment.NewPaymentService("go.micro.service.payment", srv.Client())

	// pay
	reply, err := ticketService.Pay(context.TODO(), &goMicroServicePayment.PayRequest{
		Subject:     "购票",
		OutTradeNo:  "test123456",
		TotalAmount: 0,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%v", reply)

	// Refund
	//reply, err = ticketService.Refund(context.TODO(), &goMicroServicePayment.RefundRequest{
	//
	//})
	//if err != nil {
	//	log.Fatalf("%v", err)
	//}
	//log.Printf("%v", reply)
}
