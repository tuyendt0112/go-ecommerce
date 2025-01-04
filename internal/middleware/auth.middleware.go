package middleware

import (
	"fmt"
	"go-ecommerce-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		fmt.Println("Token received:", token)
		
		if token != "valid" {
			response.ErrorResponse(ctx, response.ErrAuthentication, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
