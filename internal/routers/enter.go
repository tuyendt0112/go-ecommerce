package routers

import (
	"go-ecommerce-backend/internal/routers/manage"
	"go-ecommerce-backend/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Admin manage.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)