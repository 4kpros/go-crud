package routes

import (
	"github.com/4kpros/go-crud/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	const path = "/posts"

	r.GET(path+"/:id", controllers.GetPost)
	r.GET(path+"", controllers.GetAllPosts)
	r.POST(path+"", controllers.CreatePost)
	r.PUT(path+"/:id", controllers.UpdatePost)
	r.DELETE(path+"/:id", controllers.DeletePost)
}
