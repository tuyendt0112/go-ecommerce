package routers

import (
	"fmt"
	c "go-ecommerce-backend/internal/controller"
	"go-ecommerce-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before AA")
		c.Next()
		fmt.Println("after AA")
	}

}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before BB")
		c.Next()
		fmt.Println("after BB")
	}

}

func CC(c *gin.Context) {

	fmt.Println("before CC")
	c.Next()
	fmt.Println("after CC")

}
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthenMiddleware(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping/:name", c.NewPongController().Pong)  // Route for /v1/2024/ping/:name
		v1.GET("/user", c.NewUserController().GetUserByID) // Route for /v1/2024/pang/
	}

	return r
}
