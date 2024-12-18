package main

import (
	"house/handler"
	"house/model"
	pb "house/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "house"
	version = "latest"
)

func main() {
	model.InitDB()
	model.InitRedisPool()

	consulReg := consul.NewRegistry()
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address("127.0.0.1:12003"),
		micro.Registry(consulReg),
	)
	srv.Init()

	// Register handler
	pb.RegisterHouseHandler(srv.Server(), new(handler.House))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
