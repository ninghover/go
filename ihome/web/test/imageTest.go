package main

// import (
// "fmt"
// "image/color"
// "image/png"
// "net/http"

// "github.com/afocus/captcha"
// )

// func main() {
// 	cap := captcha.New()

// 	if err := cap.SetFont("../conf/comic.ttf"); err != nil {
// 		panic(err.Error())
// 	}

// 	cap.SetSize(128, 64)

// 	cap.SetDisturbance(captcha.NORMAL)

// 	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})

// 	cap.SetBkgColor(color.RGBA{65, 54, 45, 12}, color.RGBA{45, 120, 251, 152}, color.RGBA{102, 153, 200, 25})

// 	// 生成验证码，展示在页面中
// 	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
// 		img, str := cap.Create(4, captcha.ALL)
// 		png.Encode(w, img)
// 		fmt.Println(str)
// 	})

// 	http.ListenAndServe(":8086", nil)
// }
