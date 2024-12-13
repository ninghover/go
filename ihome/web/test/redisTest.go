package main

// import (
// 	"fmt"

// 	"github.com/gomodule/redigo/redis"
// )

// func main() {
// 	// 1.连接数据库
// 	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
// 	if err != nil {
// 		fmt.Println("redis Dial err: ", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// 2.操作数据库
// 	reply, err := conn.Do("set", "name", "huang")

// 	// 3.回复助手
// 	reply, err = redis.String(reply, err)
// 	fmt.Println(reply, err)
// }
