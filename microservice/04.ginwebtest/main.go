// 以下是一个简单的gin实例
// go run main.go 启动，占用8080端口，打印出hello gin

// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	// 1.初始化路由
// 	router := gin.Default()

// 	// 2.做路由匹配
// 	router.GET("/", func(ctx *gin.Context) {
// 		ctx.Writer.WriteString("hello gin")
// 	})

// 	// 3.启动运行
// 	router.Run(":8080")
// }

package main

import (
	"context"
	"log"
	pb "microservice/ginwebtest/proto" // microservice/ginwebtest是go.mod中的module名，proto是文件夹名
	"net/http"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func CallRemote(c *gin.Context) {
	// consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	service := micro.NewService(
		micro.Registry(consulReg),
	)
	client := pb.NewHellomicroService("hellomicro", service.Client())

	resp, err := client.Call(context.TODO(), &pb.CallRequest{Name: "huang 浩"})
	if err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "hello")
		return
	}

	c.String(http.StatusOK, resp.Msg)

}

func main() {
	// 1. 初始化路由
	router := gin.Default()

	// 2.路由匹配
	router.GET("/", CallRemote)

	// 3.run
	router.Run(":8080")
}
