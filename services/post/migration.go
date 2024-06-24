package post

import (
	"github.com/4kpros/go-crud/common/initializers"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/services/post/models"
)

func SetupMigrations() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	utils.PrintMigrationLogs(err, "Post")
}
