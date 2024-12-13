package main

import (
	"ihome/web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/",func(ctx *gin.Context) {
	// 	ctx.Writer.WriteString("web...")
	// })
	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd) //api/v1.0/imagecode/{uuid}, :uuid用于匹配

	router.GET("/api/v1.0/smscode/:phone",controller.GetSmsCd)

	router.POST("api/v1.0/users",controller.PostRet)


	router.Run(":8080")
}
