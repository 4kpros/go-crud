package main

import (
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/user"
	"go.uber.org/zap"
)

func init() {
	// Setup logger
	helpers.SetupLogger()

	// Load env variables
	errAppEnv := config.LoadAppEnv(".")
	if errAppEnv != nil {
		helpers.Logger.Warn(
			"Failed to load app ENV vars!",
			zap.String("Error", errAppEnv.Error()),
		)
		return
	}
	helpers.Logger.Warn(
		"App ENV variables loaded!",
	)

	// Connect to postgres database
	errPostgresDB := config.ConnectToPostgresDB()
	if errPostgresDB != nil {
		helpers.Logger.Warn(
			"Failed to connect to Postgres database!",
			zap.String("Error", errPostgresDB.Error()),
		)
		return
	}
	helpers.Logger.Info(
		"Connected to Postgres!",
	)
}

func main() {
	user.SetupMigrations() // User migrations
}
