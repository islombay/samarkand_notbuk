package models

import (
	"mime/multipart"
	"time"
)

type UploadFile struct {
	File *multipart.FileHeader `form:"file" binding:"required" swaggerignore:"true"`

	ID      string `json:"-" form:"-"`
	FileURL string `json:"-" form:"-"`
}

type File struct {
	ID      string `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	FileURL string `json:"file_url" gorm:"not null"`

	CreatedAt time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
