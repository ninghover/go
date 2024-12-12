package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"ihome/web/utils"
	"image/png"

	"net/http"

	getCaptcha "ihome/web/proto/getCaptcha"

	"github.com/afocus/captcha"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetSession(ctx *gin.Context) {
	res := make(map[string]string)
	res["errno"] = utils.RECODE_SESSIONERR
	res["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, res)
}

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
		fmt.Println("未找到远程服务:", err)
	}
	var img captcha.Image
	json.Unmarshal(rsp.Img, &img)

	png.Encode(ctx.Writer, &img)
}