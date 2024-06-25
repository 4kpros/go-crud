package main

import (
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
)

func init() {
	utils.InitializeLogger()
	config.LoadEnvironmentVariables(".")
	config.ConnectToDB()
}

func main() {
	post.SetupMigrations()
	auth.SetupMigrations()
}
