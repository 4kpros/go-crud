package post

import (
	"github.com/4kpros/go-crud/services/post/routes"
	"github.com/gin-gonic/gin"
)

func SetupService(r *gin.Engine) {
	routes.SetupRoutes(r)
}
