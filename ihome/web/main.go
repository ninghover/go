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

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd)

	router.Run(":8080")
}
