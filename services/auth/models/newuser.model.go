package models

import "gorm.io/gorm"

type NewUser struct {
	gorm.Model
	Email         string
	PhoneNumber   int
	Password      string
	Provider      string
	ProviderToken string
	IsValid       bool
}
