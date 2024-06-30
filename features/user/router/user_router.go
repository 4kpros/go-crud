package router

import (
	"github.com/4kpros/go-api/common/router"
	"github.com/4kpros/go-api/features/user/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserEndpoints(routerGroup *gin.RouterGroup, controller *controller.UserController) {

	group := routerGroup.Group("/users")
	const requireAuth = true

	router.POST(group, "/email", controller.CreateWithEmail, requireAuth)
	router.POST(group, "/phone", controller.CreateWithPhoneNumber, requireAuth)
	router.PUT(group, "/:id", controller.UpdateUser, requireAuth)
	router.PUT(group, "/info/:id", controller.UpdateUserInfo, requireAuth)
	router.DELETE(group, "/:id", controller.Delete, requireAuth)
	router.GET(group, "/:id", controller.FindById, requireAuth)
	router.GET(group, "", controller.FindAll, requireAuth)
}
