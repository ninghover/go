package main

import (
	"user/handler"
	"user/model"
	pb "user/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	// 初始化Redis连接池
	model.InitRedisPool()

	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
		micro.Address("127.0.0.1:12001"),
	)
	srv.Init()

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
