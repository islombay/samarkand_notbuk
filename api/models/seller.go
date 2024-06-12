package models

import "time"

type Seller struct {
	ID          string `gorm:"primaryKey" json:"id,omitempty"`
	FirstName   string `gorm:"size:40" json:"first_name,omitempty"`
	LastName    string `gorm:"size:40" json:"last_name,omitempty"`
	PhoneNumber string `gorm:"size:12;unique" json:"phone_number,omitempty"`

	CreatedAt time.Time  `gorm:"autoCreateTime:milli" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateSeller struct {
	FirstName   string `json:"first_name" binding:"required,max=40"`
	LastName    string `json:"last_name" binding:"required,max=40"`
	PhoneNumber string `json:"phone_number" binding:"required,max=12"`
}

type UpdateSeller struct {
	ID          string `json:"id" binding:"required"`
	FirstName   string `json:"first_name" binding:"omitempty,max=40"`
	LastName    string `json:"last_name" binding:"omitempty,max=40"`
	PhoneNumber string `json:"phone_number"`
}
