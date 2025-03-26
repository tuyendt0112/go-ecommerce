package manage

import "github.com/gin-gonic/gin"

type AdminUserRouter struct {
}

func (ar *AdminUserRouter) InitAdminUserRouter(Router *gin.RouterGroup) {

	// private
	adminRouterPrivate := Router.Group("admin/user")
	{
		adminRouterPrivate.GET("/user/list")
	}
}
