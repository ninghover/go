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

	// 中间件，对该中间件之前的路由不生效，对该中间件之后的路由全部生效
	router.Use(sessions.Sessions("login", store))

	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)

	router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd) //api/v1.0/imagecode/{uuid}, :uuid用于匹配

	router.GET("/api/v1.0/smscode/:phone", controller.GetSmsCd)

	router.POST("api/v1.0/users", controller.PostRet)

	router.GET("/api/v1.0/areas", controller.GetArea)

	router.POST("/api/v1.0/sessions", controller.PostLogin)

	// 校验是否登录，没登陆就不能执行下面的
	router.Use(controller.Filter)

	router.DELETE("/api/v1.0/session", controller.DeleteSession)

	router.GET("/api/v1.0/user", controller.GetUserInfo)

	router.PUT("/api/v1.0/user/name", controller.PutUserName) // 改

	// 文件上传
	router.POST("/api/v1.0/user/avatar", controller.Avatar)

	// 实名认证
	router.POST("/api/v1.0/user/auth", controller.PostUserAuth) // 增

	// 获取实名认证信息
	router.GET("/api/v1.0/user/auth", controller.GetUserInfo)

	// 获取用户已发布房源
	router.GET("api/v1.0/user/houses", controller.GetUserHouses)

	// 用户发布房源
	router.POST("/api/v1.0/houses", controller.PostHouses)

	router.Run(":8080")
}
