package main

import (
	"context"
	"errors"
	"fmt"
	pb "microservice/grpctest/protobuf/goods" //go mod中的模块名
	"net"

	"google.golang.org/grpc"
)

type Goods struct {
	pb.UnimplementedGoodsRpcServer
}

// 重写方法
func (g *Goods) GetGoods(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	var (
		name string
		err  error
	)
	if req.Id == 0 {
		err = errors.New("商品不存在")
	} else {
		name = fmt.Sprintf("%d号商品", req.Id)
	}
	return &pb.Response{
		Name:  name,
		Price: 20,
	}, err
}

func main() {
	listen, err := net.Listen("tcp", ":9527") // 启动一个tcp监听器
	if err != nil {
		fmt.Println("监听失败")
		return
	}
	server := grpc.NewServer()                  // 创建一个新的grpc服务器实例
	pb.RegisterGoodsRpcServer(server, &Goods{}) // 将Goods实例注册到grpc服务器
	err = server.Serve(listen)                  // 启动服务器 （服务器会阻塞，直到服务停止或错误）
	if err != nil {
		fmt.Println("启动服务失败")
	}
}
