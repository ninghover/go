package main

import (
	"hellomicro/handler"
	pb "hellomicro/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"	// add
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"	// add
)

var (
	service = "hellomicro"
	version = "latest"
)

func main() {
	//Register cunsul
	consulReg:=consul.NewRegistry(func (options *registry.Options)  { // add
		options.Addrs=[]string{"127.0.0.1:8500"}	// add
	})	// add


	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),	// add
	)
	srv.Init()

	// Register handler
	pb.RegisterHellomicroHandler(srv.Server(), new(handler.Hellomicro))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
