package auth

import (
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/auth/model"
)

func SetupMigrations() {
	err := config.DB.AutoMigrate(&model.NewUser{})
	utils.PrintMigrationLogs(err, "NewUser")
}
