package models

import (
	"time"

	"github.com/4kpros/go-crud/common/types"
	"github.com/4kpros/go-crud/common/utils"
	"gorm.io/gorm"
)

type NewUser struct {
	types.BaseGormModel
	Email          string     `json:"email"`
	PhoneNumber    int        `json:"phoneNumber"`
	Provider       string     `json:"provider"`
	ProviderUserId string     `json:"providerUserId"`
	IsValid        bool       `json:"isValid"`
	ValidatedAt    *time.Time `json:"validatedAt"`
	Password       string     `json:"password"`
}

func (u *NewUser) BeforeCreate(db *gorm.DB) (err error) {
	u.Password, err = utils.EncryptWithArgon2id(u.Password)
	return
}
