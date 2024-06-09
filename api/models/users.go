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
