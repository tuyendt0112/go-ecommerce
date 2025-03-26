//go:build wireinject

package wire

import (
	"go-ecommerce-backend/internal/controller"
	"go-ecommerce-backend/internal/repo"
	"go-ecommerce-backend/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
