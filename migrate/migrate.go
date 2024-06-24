package main

import (
	"github.com/4kpros/go-crud/common/initializers"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/services/auth"
	"github.com/4kpros/go-crud/services/post"
)

func init() {
	utils.InitializeLogger()
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	post.SetupMigrations()
	auth.SetupMigrations()
}
