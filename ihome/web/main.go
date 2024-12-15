package main

import (
	"ihome/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化redis容器，存放session
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("hello"))
	router.Use(sessions.Sessions("login", store))

	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd) //api/v1.0/imagecode/{uuid}, :uuid用于匹配

	router.GET("/api/v1.0/smscode/:phone", controller.GetSmsCd)

	router.POST("api/v1.0/users", controller.PostRet)

	router.GET("/api/v1.0/areas", controller.GetArea)

	router.POST("/api/v1.0/sessions", controller.PostLogin)


	router.DELETE("/api/v1.0/session",controller.DeleteSession)

	router.Run(":8080")
}
