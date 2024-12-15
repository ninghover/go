package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"ihome/web/utils"
	"image/png"

	"net/http"

	getArea "ihome/web/proto/getArea"
	getCaptcha "ihome/web/proto/getCaptcha"
	user "ihome/web/proto/user"

	"github.com/afocus/captcha"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetSession(ctx *gin.Context) {
	res := make(map[string]interface{})
	session := sessions.Default(ctx)
	name := session.Get("name")
	fmt.Println("name=", name)
	if name == nil {
		res["errno"] = utils.RECODE_SESSIONERR
		res["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		res["errno"] = utils.RECODE_OK
		res["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		var logData struct {
			Name string `json:"name"`
		}
		logData.Name = name.(string)
		res["data"] = logData
	}
	ctx.JSON(http.StatusOK, res)
}

// 获取图片验证码
func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	consulReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Registry(consulReg),
		// micro.Address("127.0.0.1:8500"), //如果未显式指定 micro.Address，micro.NewService 会使用默认的注册中心地址(通常是 127.0.0.1:8500)
	)
	client := getCaptcha.NewGetCaptchaService("getCaptcha", srv.Client())
	rsp, err := client.Call(context.TODO(), &getCaptcha.CallRequest{Uuid: uuid})
	if err != nil {
		fmt.Println("GetImageCd 远程调用失败:", err)
	}
	var img captcha.Image
	json.Unmarshal(rsp.Img, &img)

	png.Encode(ctx.Writer, &img)
}

// 获取短信验证码
func GetSmsCd(ctx *gin.Context) {
	phone := ctx.Param("phone")
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	consulReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	client := user.NewUserService("user", srv.Client())
	rsp, err := client.SendSms(context.TODO(), &user.SmsReq{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("GetSmsCd 远程调用失败:", err)
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 注册新用户
func PostRet(ctx *gin.Context) {
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	ctx.Bind(&regData)
	consulReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	client := user.NewUserService("user", srv.Client())
	rsp, err := client.Register(context.TODO(), &user.RegReq{Mobile: regData.Mobile, Password: regData.PassWord, Smscode: regData.SmsCode})
	fmt.Println("rsp=", rsp)
	fmt.Println("err=", err)
	if err != nil {
		fmt.Println("PostRet 远程调用失败")
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 获取城区信息
func GetArea(ctx *gin.Context) {
	consulReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	client := getArea.NewGetAreaService("getarea", srv.Client())
	rsp, err := client.GetArea(context.TODO(), &getArea.AreaReq{})
	if err != nil {
		fmt.Println("GetArea 远程调用失败")
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 发送登录信息服务
func PostLogin(ctx *gin.Context) {
	var logData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	ctx.Bind(&logData)
	consulReg := consul.NewRegistry()
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	client := user.NewUserService("user", srv.Client())
	rsp, _ := client.Login(ctx, &user.LoginReq{Mobile: logData.Mobile, Password: logData.Password})
	if rsp.Errno != utils.RECODE_OK {
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	session := sessions.Default(ctx)
	session.Set("name", rsp.Name)
	session.Save()
	ctx.JSON(http.StatusOK, rsp)
}

func DeleteSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	rsp := make(map[string]string)
	session.Delete("name")
	if err := session.Save(); err != nil {
		rsp["errno"] = utils.RECODE_DATAERR
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	} else {
		rsp["errno"] = utils.RECODE_OK
		rsp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, rsp)
}
