package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware gin.HandlerFunc

func NewAuthMiddleware() AuthMiddleware {
	return func(ctx *gin.Context) {
		// TODO: implement authentication logic
		ctx.Next()
	}
}
