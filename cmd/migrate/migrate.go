package migrate

import (
	"github.com/4kpros/go-api/features/user"
)

func Start() {
	user.SetupMigrations() // User migrations
}
