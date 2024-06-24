package routes

import (
	"github.com/4kpros/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	const path = "/posts"

	GET(r, path+"/:id", controllers.GetPost, true)
	GET(r, path+"", controllers.GetAllPosts, true)
	POST(r, path+"", controllers.CreatePost, true)
	PUT(r, path+"/:id", controllers.UpdatePost, true)
	PATCH(r, path+"/:id", controllers.UpdatePost, true)
	DELETE(r, path+"/:id", controllers.DeletePost, true)
}
