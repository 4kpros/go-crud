package main

import (
	"github.com/4kpros/go-crud/initializers"
	"github.com/4kpros/go-crud/migrations"
	"github.com/4kpros/go-crud/utils"
)

func init() {
	utils.InitializeLogger()
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	migrations.PostMigrations()
}
