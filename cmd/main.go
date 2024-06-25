package main

import (
	"github.com/4kpros/go-crud/common/middlewares"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitializeLogger()
	config.LoadEnvironmentVariables(".")
	config.ConnectToDB()
}

func main() {
	// Setup gin for HTTP requests
	r := gin.Default()
	r.Use(middlewares.ErrorsHandler())

	// Setup endpoints
	auth.SetupService(r)
	post.SetupService(r)

	// Run gin
	r.Run(":" + config.EnvConfig.ServerPort)
}
