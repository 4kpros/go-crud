package initializers

import (
	"github.com/4kpros/go-crud/utils"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Warn(
			"Failed to load ENV vars !",
			zap.String("Error", err.Error()),
		)
	}
}
