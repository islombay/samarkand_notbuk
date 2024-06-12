package models

import (
	"time"
)

type Client struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string `gorm:"size:40"`
	LastName    string `gorm:"size:40"`
	PhoneNumber string `gorm:"size:12;unique"`
	Password    string

	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time
}

type Staff struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string `gorm:"size:40"`
	LastName    string `gorm:"size:40"`
	PhoneNumber string `gorm:"size:12;unique"`
	Password    string

	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time
}

type UpdateClient struct {
	ID        string `js:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"omitempty,max=40"`
	LastName  string `json:"last_name" binding:"omitempty,max=40"`
}
