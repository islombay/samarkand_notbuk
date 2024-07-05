package models

import (
	"time"
)

type Product struct {
	ID          string  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	Price       float64 `gorm:"type:numeric;not null" json:"price,omitempty"`

	CategoryID *string `gorm:"type:uuid" json:"category_id,omitempty"`
	BrandID    *string `gorm:"type:uuid" json:"brand_id,omitempty"`
	ImageID    *string `gorm:"type:uuid" json:"image_id,omitempty"`
	VideoID    *string `gorm:"type:uuid" json:"video_id,omitempty"`

	CreatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`

	Category *Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"category,omitempty"`
	Brand    *Brand    `gorm:"foreignKey:BrandID;constraint:OnDelete:SET NULL" json:"brand,omitempty"`
	Image    *File     `gorm:"foreignKey:ImageID;constraint:OnDelete:SET NULL" json:"image,omitempty"`
	Video    *File     `gorm:"foreignKey:VideoID;constraint:OnDelete:SET NULL" json:"video,omitempty"`

	ProductFiles        []ProductFile        `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"files,omitempty"`
	ProductInstallments []ProductInstallment `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"installments,omitempty"`
}

type ProductInstallment struct {
	ProductID string  `gorm:"type:uuid;not null" json:"product_id"`
	Price     float64 `gorm:"type:numeric;not null" json:"price"`
	Period    int     `gorm:"type:int;not null" json:"period"`

	CreatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`

	//Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
}

type ProductFile struct {
	ProductID string `gorm:"type:uuid;not null" json:"product_id,omitempty"`
	FileID    string `gorm:"type:uuid;not null" json:"file_id,omitempty"`

	CreatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`

	// Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
	// File    File    `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty"`
}

type CreateProduct struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" binding:"required"`

	CategoryID *string `json:"category_id" binding:"omitempty,uuid"`
	BrandID    *string `json:"brand_id" binding:"omitempty,uuid"`
	ImageID    *string `json:"image_id" binding:"omitempty,uuid"`
	VideoID    *string `json:"video_id" binding:"omitempty,uuid"`

	ProductFiles        []CreateProductFiles        `json:"files"`
	ProductInstallments []CreateProductInstallments `json:"installments"`
}

type CreateProductFiles struct {
	FileID string `json:"file_id" binding:"required,uuid"`
}

type CreateProductInstallments struct {
	Price  float64 `json:"price" binding:"required"`
	Period int     `json:"period" binding:"required"`
}
