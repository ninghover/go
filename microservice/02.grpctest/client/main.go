package main

import (
	"context"
	"fmt"
	"time"

	pb "microservice/grpctest/protobuf/goods"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9527", grpc.WithTransportCredentials(insecure.NewCredentials())) // 创建一个连接，不安全的连接
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	defer conn.Close()
	c := pb.NewGoodsRpcClient(conn)                                       // 创建一个GoodsRpcClient类型的gRPC客户端，conn 是连接对象
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // 创建一个新的contexe，客户端请求必须在1s内返回
	defer cancel()                                                        // 取消上下文，释放资源
	res, err := c.GetGoods(ctx, &pb.Request{Id: 5})
	if err != nil {
		fmt.Println("无法调用")
		return
	}
	fmt.Println("商品名：", res.GetName())
	fmt.Println("价格：", res.GetPrice())
}
