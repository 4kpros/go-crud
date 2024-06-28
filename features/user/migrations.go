package user

import (
	"github.com/4kpros/go-api/common/helpers"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/user/model"
)

func SetupMigrations() {
	err := config.DB.AutoMigrate(&model.User{})
	helpers.PrintMigrationLogs(err, "User")
}
