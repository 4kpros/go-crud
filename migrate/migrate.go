package main

import (
	"github.com/4kpros/go-crud/initializers"
	"github.com/4kpros/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
