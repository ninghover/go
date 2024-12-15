package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func readFile(path string) (id, secret string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line, "ID:") {
			id = strings.TrimSpace(strings.TrimPrefix(line, "ID:"))
		} else if strings.HasPrefix(line, "Secret:") {
			secret = strings.TrimSpace(strings.TrimPrefix(line, "Secret:"))
		}
	}
	return
}

func CreateClient() (*dysmsapi20170525.Client, error) {
	id, secret, err := readFile("../../secret.txt") // 相对于启动路径的
	if err != nil {
		return nil, err
	}
	fmt.Println(id, secret)
	config := &openapi.Config{
		AccessKeyId:     tea.String(id),
		AccessKeySecret: tea.String(secret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	return dysmsapi20170525.NewClient(config)
}

func main() {
	client, err := CreateClient()
	if err != nil {
		panic(err)
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String("15828977067"),
		SignName:      tea.String("go语言租房网"),
		TemplateCode:  tea.String("SMS_476270071"),
		TemplateParam: tea.String(`{"code":"567845"}`), // 4-6位验证码
	}

	runtime := &util.RuntimeOptions{}
	_, err = client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		panic(err)
	}
}

// https://next.api.aliyun.com/api/Dysmsapi/2017-05-25/SendSms?params={%22SignName%22:%22go%E8%AF%AD%E8%A8%80%E7%A7%9F%E6%88%BF%E7%BD%91%22,%22TemplateCode%22:%22SMS_476270071%22,%22PhoneNumbers%22:%2215828977067%22,%22TemplateParam%22:%22%7B%5C%22code%5C%22%3A%5C%221234%5C%22%7D%22}&RegionId=cn-qingdao&tab=DEMO&lang=GO
