package user

import (
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/user/model"
)

func SetupMigrations() {
	err := config.DB.AutoMigrate(&model.User{})
	utils.PrintMigrationLogs(err, "User")
}
