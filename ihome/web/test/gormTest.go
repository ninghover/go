// 需要安装的驱动
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func main() {
// 	// 配置数据库连接信息
// 	// dsn := "root:000914@tcp(127.0.0.1:3306)/ihome?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "root:000914@tcp(127.0.0.1:3306)/ihome"

// 	// 初始化数据库连接
// 	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	fmt.Println(conn)
// 	if err != nil {
// 		log.Fatal("数据库连接失败:", err)
// 	}

// 	log.Println("数据库连接成功!")
// }
