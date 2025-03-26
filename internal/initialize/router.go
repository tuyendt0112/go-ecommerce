package initialize

import (
	"go-ecommerce-backend/global"
	"go-ecommerce-backend/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	r.Use() // logger
	r.Use() // cross
	r.Use() // limiter
	userRouter := routers.RouterGroupApp.User
	adminRouter := routers.RouterGroupApp.Admin

	MainGroup := r.Group("v1/2024")
	{
		MainGroup.GET("/check")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
		adminRouter.InitAdminUserRouter(MainGroup)
	}
	return r
}
