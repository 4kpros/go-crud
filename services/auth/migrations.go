package auth

import (
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth/models"
)

func SetupMigrations() {
	err := config.DB.AutoMigrate(&models.NewUser{})
	utils.PrintMigrationLogs(err, "NewUser")
}
