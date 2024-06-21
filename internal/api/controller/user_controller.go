package controllers

import (
	"A2BackEnd/internal/repositories"
	"A2BackEnd/internal/services"
	"A2BackEnd/internal/utils"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return &UserController{userService}
}

func (c *UserController) GetCaptcha(ctx iris.Context) {
	captchaID := captcha.New()
	ctx.JSON(iris.Map{
		"captcha_id":  captchaID,
		"captcha_url": "/captcha/" + captchaID + ".png",
	})
}

func (c *UserController) Register(ctx iris.Context) {
	var req struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		CaptchaID string `json:"captcha_id"`
		Captcha   string `json:"captcha"`
	}
	if err := ctx.ReadJSON(&req); err != nil {
		utils.JSONResponse(ctx, 400, "Invalid request", nil)
		return
	}
	if !captcha.VerifyString(req.CaptchaID, req.Captcha) {
		utils.JSONResponse(ctx, 400, "Invalid captcha", nil)
		return
	}
	if err := c.userService.RegisterUser(req.Username, req.Password); err != nil {
		utils.JSONResponse(ctx, 500, "Could not register user", nil)
		return
	}
	utils.JSONResponse(ctx, 200, "User registered successfully", nil)
}

func (c *UserController) Login(ctx iris.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ReadJSON(&req); err != nil {
		utils.JSONResponse(ctx, 400, "Invalid request", nil)
		return
	}
	token, err := c.userService.LoginUser(req.Username, req.Password)
	if err != nil {
		utils.JSONResponse(ctx, 401, "Invalid credentials", nil)
		return
	}
	utils.JSONResponse(ctx, 200, "Login successful", map[string]string{"token": token})
}
