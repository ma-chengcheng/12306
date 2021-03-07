package main

import (
	"github.com/mamachengcheng/12306/common"
	"github.com/mamachengcheng/12306/srv/user/domain/respository"
	s "github.com/mamachengcheng/12306/srv/user/domain/service"
	"github.com/mamachengcheng/12306/srv/user/handler"
	user2 "github.com/mamachengcheng/12306/srv/user/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func main() {
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")

	if err != nil {
		log.Fatalf("%v", err)
	}

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	dsn := mysqlInfo.User + ":" + mysqlInfo.Password + "@tcp(" + mysqlInfo.Host + ":" + strconv.FormatInt(mysqlInfo.Port, 10) + ")/" + mysqlInfo.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 缓存预编译语句
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	//respository.NewUserRepository(db).InitTable()

	userDataService := s.NewUserDataService(respository.NewUserRepository(db))

	// Register handler
	user2.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
