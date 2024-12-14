package handler

import (
	"context"
	"encoding/json"
	"getCaptcha/model"
	pb "getCaptcha/proto"
	"image/color"

	"github.com/afocus/captcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	cap := captcha.New()

	if err := cap.SetFont("./conf/comic.ttf"); err != nil { // 这个路径是相对于运行路径的，对于main.go来说 是这个路径
		panic(err.Error())
	}

	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{65, 54, 45, 12}, color.RGBA{45, 120, 251, 152}, color.RGBA{102, 153, 200, 25})
	img, code := cap.Create(4, captcha.NUM)

	// 将验证码保存到redis
	model.SaveImgCode(req.GetUuid(), code)

	rsp.Img, _ = json.Marshal(img)
	return nil
}
