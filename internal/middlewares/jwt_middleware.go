package middlewares

import (
	"A2BackEnd/internal/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
	"strings"
)

func JWTMiddleware(ctx iris.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		utils.JSONResponse(ctx, 401, "Authorization header is missing", nil)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		utils.JSONResponse(ctx, 401, "Invalid Authorization header format", nil)
		return
	}

	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // Replace "secret" with your actual secret key
	})

	if err != nil || !token.Valid {
		utils.JSONResponse(ctx, 401, "Invalid token", nil)
		return
	}

	ctx.Next()
}
