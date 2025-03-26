package user

import (
	"go-ecommerce-backend/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	// public
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/otp")
		userRouterPublic.POST("/register", userController.Register)

	}

	// private

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.User(Limiter())
	{
		userRouterPrivate.GET("/info")

	}
}
