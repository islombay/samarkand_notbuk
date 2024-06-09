package models

import "time"

type Category struct {
	ID       string  `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	NameUz   string  `gorm:"size:30;not null" json:"name"`
	ParentID *string `json:"parent_id,omitempty"`

	Parent        *Category  `gorm:"foreignKey:parent_id;references:id" json:"parent,omitempty"`
	Subcategories []Category `gorm:"foreignKey:parent_id;references:id" json:"subcategories,omitempty"`

	CreatedAt time.Time  `gorm:"autoCreateTime:milli" json:"created_at,omitempty"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CreateCategory struct {
	NameUz   string  `json:"name" binding:"required,max=30"`
	ParentID *string `json:"parent_id"`
}

type UpdateCategory struct {
	ID       string  `json:"id" binding:"required"`
	NameUz   *string `json:"name" binding:"omitempty,max=30"`
	ParentID *string `json:"parent_id"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required"`
}
