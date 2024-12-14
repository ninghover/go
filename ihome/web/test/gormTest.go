// 需要安装的驱动
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql
// 参考博客 https://www.cnblogs.com/smartljy/p/18471269

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 2.定义结构体模型
type Student struct {
	// gorm.Model // 软删除时可以使用
	Id   uint   // 默认主键，自增
	Name string // 如果是私有的就表中就不会创建
	Age  uint
}

// type User struct {
// 	// gorm.Model // 软删除时可以使用
// 	Id   uint
// 	Name string `gorm:"size:100;default:'xiaoming'"`
// 	Age  uint   `gorm:"not null"`
// }

func main() {
	// 1. 连接数据库
	// 配置数据库连接信息
	dsn := "root:000914@tcp(127.0.0.1:3306)/ihome?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:000914@tcp(127.0.0.1:3306)/ihome"

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	log.Println("数据库连接成功!", db)

	// 2.根据结构体生成数据库表
	db.AutoMigrate(&Student{})

	// 3.使用Create新增数据
	s := Student{Name: "zhangsan", Age: 18}
	if res := db.Create(&s); res.Error != nil {
		log.Println("插入失败!", res)
	} else {

		log.Println("插入成功!", res)
	}

	// 4.查询
	// var s Student
	// db.First(&s,2)  //  默认会根据表中的主键查询
	// db.Find(&s,"age>?","20")

	var ss []Student
	db.Where("age>? and name =?", 20, "huang").Find(&ss) // db.Find(&ss) 直接查全部
	fmt.Println(ss)

	// 5.更新
	// s:=Student{Id:1}
	// db.Model(&s).Update("age",30)    // 更新单个字段，id为3的，age改为30

	// db.Model(&s).Updates(&Student{Name: "黄",Age: 24})  // 更新多个字段，id为1的

	// db.Where("age = ?", 30).Model(&Student{}).Update("name", "update") // 根据指定条件来更新

	// 6.删除
	// 根据主键删除
	// db.Delete(&Student{}, 1)

	// 根据条件删除
	// db.Where("age = ?", 18).Delete(&Student{})

}
