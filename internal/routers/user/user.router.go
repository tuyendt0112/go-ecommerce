package user

import "github.com/gin-gonic/gin"

type UserRouter struct {

}

func (ur *UserRouter) InitUserRouter (Router *gin.RouterGroup){

	// public 
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/otp")
		userRouterPublic.POST("/register")
		
	}



	// private

	userRouterPrivate :=Router.Group("/user")
	// userRouterPrivate.User(Limiter())
	{
		userRouterPrivate.GET("/info")
		
	}
}
