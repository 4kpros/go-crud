package main

import (
	"github.com/4kpros/go-crud/common/middlewares"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// Setup logger
	utils.InitializeLogger()

	// Load env variables
	errEnv := config.LoadEnvironmentVariables(".")
	if errEnv != nil {
		utils.Logger.Warn(
			"Failed to load ENV vars !",
			zap.String("Error", errEnv.Error()),
		)
		return
	}
	utils.Logger.Warn(
		"Env variables loaded !",
	)

	// Connect to postgres database
	errPostgresDB := config.ConnectToPostgresDB()
	if errPostgresDB != nil {
		utils.Logger.Warn(
			"Failed to connect to Postgres database !",
			zap.String("Error", errPostgresDB.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Connected to Postgres database: ",
		zap.String("DB name", config.DB.Name()),
	)
}

func main() {
	// Setup gin for HTTP requests
	r := gin.Default()
	r.Use(middlewares.ErrorsHandler())

	// Setup services
	auth.SetupService(r) // Auth service
	post.SetupService(r) // Post service

	// Run gin with custom port
	r.Run(":" + config.EnvConfig.ServerPort)
}
