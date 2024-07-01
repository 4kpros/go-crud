package migrate

import (
	userMigrate "github.com/4kpros/go-api/services/user/migrate"
)

func Start() {
	userMigrate.Migrate() // User migrations
}
