package main

import (
	controllers "A2BackEnd/internal/api/controller"
	"A2BackEnd/internal/config"
	"A2BackEnd/internal/middlewares"
	"A2BackEnd/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {
	app := iris.New()

	// Database connection
	db := config.SetupDatabase()

	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)

	app.Use(middlewares.Cors)

	// User routes
	apiRouter := app.Party("/v1")
	{
		userController := controllers.NewUserController(db)
		userRouter := apiRouter.Party("/users")
		{
			userRouter.Get("/captcha", userController.GetCaptcha)
			userRouter.Get("/captcha/{captchaID:string}", utils.CaptchaServe)
			userRouter.Post("/register", userController.Register)
			userRouter.Post("/login", userController.Login)
		}
	}

	// Protected routes
	app.Use(middlewares.JWTMiddleware)
	// Add protected routes here

	app.Listen(":8080")
}
