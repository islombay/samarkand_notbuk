package models

import (
	"time"
)

type Brand struct {
	ID   string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string `gorm:"size:40" json:"name"`

	CreatedAt time.Time  `gorm:"autoCreateTime:milli" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateBrand struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBrand struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
