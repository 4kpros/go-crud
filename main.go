package main

import (
	"github.com/4kpros/go-crud/initializers"
	"github.com/4kpros/go-crud/middleware"
	"github.com/4kpros/go-crud/routes"
	"github.com/4kpros/go-crud/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.InitializeLogger()
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	//Setup gin
	r := gin.Default()
	r.Use(middleware.ErrorsHandler())

	// Setup routes
	routes.PostRoutes(r)

	// Listen
	r.Run()
}
