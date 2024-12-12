package main

import (
	"getCaptcha/handler"
	pb "getCaptcha/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	// "go-micro.dev/v4/registry"
)

var (
	service = "getCaptcha"
	version = "latest"
)

func main() {
	// consulReg := consul.NewRegistry(func(o *registry.Options) {
	// 	o.Addrs = []string{"127.0.0.1:8500"}
	// })

	consulReg := consul.NewRegistry() // 可以不指定地址，默认就是127.0.0.1:8500

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
		micro.Address("127.0.0.1:12000"), // 本服务默认运行在12000端口，如果没指定，会随机分配
	)
	srv.Init()

	// Register handler
	pb.RegisterGetCaptchaHandler(srv.Server(), new(handler.GetCaptcha))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
