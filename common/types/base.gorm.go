package types

import (
	"time"

	"gorm.io/gorm"
)

type BaseGormModel struct {
	ID        uint           `gorm:"primaryKey; not null" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
