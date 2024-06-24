package routes

import (
	"github.com/4kpros/go-crud/common/router"
	"github.com/4kpros/go-crud/services/post/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	const path = "/posts"

	router.GET(r, path+"/:id", controllers.GetPost, true)
	router.GET(r, path+"", controllers.GetAllPosts, true)
	router.POST(r, path+"", controllers.CreatePost, true)
	router.PUT(r, path+"/:id", controllers.UpdatePost, true)
	router.PATCH(r, path+"/:id", controllers.UpdatePost, true)
	router.DELETE(r, path+"/:id", controllers.DeletePost, true)
}
