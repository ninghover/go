package main

import (
	"getArea/handler"
	"getArea/model"
	pb "getArea/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getarea"
	version = "latest"
)

func main() {
	model.InitRedisPool()
	model.InitDB()

	// Create service
	consulReg := consul.NewRegistry()

	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
		micro.Address("127.0.0.1:12002"),
	)
	srv.Init()

	// Register handler
	pb.RegisterGetAreaHandler(srv.Server(), new(handler.GetArea))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
