package utils

import (
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"strings"
)

func CaptchaServe(ctx iris.Context) {
	captchaID := ctx.Params().Get("captchaID")
	if strings.Contains(captchaID, ".png") {
		captchaID = strings.ReplaceAll(captchaID, ".png", "")
	}
	if captchaID == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Captcha ID is missing")
		return
	}
	ctx.ResponseWriter().Header().Set("Content-Type", "image/png")
	if err := captcha.WriteImage(ctx.ResponseWriter(), captchaID, captcha.StdWidth, captcha.StdHeight); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Failed to generate captcha image")
	}
}
