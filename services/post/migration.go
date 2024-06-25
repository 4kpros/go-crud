package post

import (
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/post/models"
)

func SetupMigrations() {
	err := config.DB.AutoMigrate(&models.Post{})
	utils.PrintMigrationLogs(err, "Post")
}
