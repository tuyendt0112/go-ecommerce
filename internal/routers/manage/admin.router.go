package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private
	adminRouterPrivate := Router.Group("/admin/product")
	{
		adminRouterPrivate.POST("/create")
	}
}
