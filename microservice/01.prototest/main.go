package main

import (
	"fmt"
	prototest "microservice/prototest/github.com/hello/prototest"
	"google.golang.org/protobuf/proto"
)

func main() {
	user := &prototest.User{
		Name:  "huang hao",
		Age:   18,
		Hobby: []string{"唱", "跳", "rap", "篮球"},
	}
	fmt.Println("user: ", user)

	data, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("编码失败")
		return
	}
	fmt.Println("编码: ", data)
	fmt.Println("字符串转换: ", string(data))

	newUser := &prototest.User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		fmt.Println("解码失败")
		return
	}
	fmt.Println("newUser: ", newUser)
	fmt.Println("----------------")
	newUser.Name = "张三"
	fmt.Println(newUser.GetName())
}

// 直接go run main.go 会说依赖找不到
// 1. go mod init
// 2. go mod tidy
// 参考 https://offernow.cn/s/language/golang/best_practice/mgkiu0x8upweqrsb
