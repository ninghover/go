package main

import (
	"fmt"

	"github.com/tedcy/fdfs_client"
)

func main() {
	clt, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Println("初始化客户端错误, err: ", err)
	}

	// 上传文件
	rsp, err := clt.UploadByFilename("01.jpg")
	fmt.Println(rsp, err)
}
