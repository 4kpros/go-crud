package migrations

import (
	"github.com/4kpros/go-crud/initializers"
	"github.com/4kpros/go-crud/models"
	"github.com/4kpros/go-crud/utils"
	"go.uber.org/zap"
)

func PostMigrations() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		utils.Logger.Warn(
			"Failed to create << Post >> table !",
			zap.String("Error", err.Error()),
		)
		return
	}
	utils.Logger.Info(
		"Migration for << Post >> Done !",
	)
}
