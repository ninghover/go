package handler

import (
	"context"
	"fmt"
	"math/rand"
	"user/model"
	pb "user/proto"
	"user/utils"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type User struct{}

func (e *User) SendSms(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	// req.ImgCode
	// req.Phone
	// req.Uuid

	// 图片验证码校验失败
	if req.ImgCode != model.GetImgCode(req.Uuid) {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}
	// 发送短信验证码
	client, _ := func() (*dysmsapi20170525.Client, error) {
		id, secret, err := utils.ReadIdAndSecret("../../secret.txt") //相对于启动路径的地址（main.go）
		fmt.Println(id, secret, err)
		if err != nil {
			return nil, err
		}
		config := &openapi.Config{
			AccessKeyId:     tea.String(id),
			AccessKeySecret: tea.String(secret),
			Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		}
		return dysmsapi20170525.NewClient(config)
	}()
	smsCode := fmt.Sprintf("%04d", rand.Int31()%10000)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(req.Phone),
		SignName:      tea.String("go语言租房网"),
		TemplateCode:  tea.String("SMS_476270071"),
		TemplateParam: tea.String(`{"code":"` + smsCode + `"}`), // 4位验证码
	}
	runtime := &util.RuntimeOptions{}
	_, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil { //发送验证码失败
		rsp.Errno = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		return err
	}

	// 将短信验证码写入redis
	model.SaveSms(req.Phone, smsCode)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}
