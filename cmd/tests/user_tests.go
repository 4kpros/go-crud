package tests

import (
	"testing"

	"github.com/4kpros/go-api/services/user"
)

func UserTests(t *testing.T) {
	repo := user.NewUserRepositoryImpl(nil)
	service := user.NewUserServiceImpl(repo)
	controller := user.NewUserController(service)

	if repo == nil || service == nil || controller == nil {
		return
	}

	// TODO add tests
}
