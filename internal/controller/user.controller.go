package controller

import (
	"go-ecommerce-backend/internal/service"
	"go-ecommerce-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service-> repo->modal->db
func (uc *UserController) GetUserByID(ctx *gin.Context) {

	response.SuccessResponse(ctx, 20001, "hello jezz")
}
