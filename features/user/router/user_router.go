package router

import (
	"github.com/4kpros/go-api/common/router"
	"github.com/4kpros/go-api/features/user/controller"
	"github.com/gin-gonic/gin"
)

func SetupUserEndpoints(routerGroup *gin.RouterGroup, controller *controller.UserController) {

	group := routerGroup.Group("/users")
	const requireAuth = true

	router.POST(group, "", controller.Create, requireAuth)       // Create
	router.PUT(group, "/:id", controller.Update, requireAuth)    // Update
	router.DELETE(group, "/:id", controller.Delete, requireAuth) // Delete
	router.GET(group, "/:id", controller.FindById, requireAuth)  // Find by id
	router.GET(group, "", controller.FindAll, requireAuth)       // Find all
}
