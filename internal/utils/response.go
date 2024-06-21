package utils

import (
	"github.com/kataras/iris/v12"
)

func JSONResponse(ctx iris.Context, status int, message string, data interface{}) {
	ctx.StatusCode(status)
	ctx.JSON(iris.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
