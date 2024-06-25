package main

import (
	"github.com/4kpros/go-crud/common/initializers"
	"github.com/4kpros/go-crud/common/middlewares"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitializeLogger()
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	// Setup gin for HTTP requests
	r := gin.Default()
	r.Use(middlewares.ErrorsHandler())

	// Setup endpoints
	auth.SetupService(r)
	post.SetupService(r)

	// Run gin
	r.Run()
}
