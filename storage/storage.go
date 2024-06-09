package storage

import (
	"context"
	"fmt"

	"github.com/islombay/noutbuk_seller/api/models"
)

var (
	ErrNotFound      = fmt.Errorf("err_not_found")
	ErrAlreadyExists = fmt.Errorf("already_exists")
)

type StorageInterface interface {
	Close()
	User() UserI
	Category() CategoryI
	Brand() BrandI
}

type UserI interface {
	GetStaffByPhoneNumber(context.Context, string) (*models.Staff, error)
}

type CategoryI interface {
	Create(context.Context, models.CreateCategory) (*models.Category, error)
	GetList(context.Context, models.Pagination) (*models.Pagination, error)
	GetByID(context.Context, string) (*models.Category, error)
	Update(ctx context.Context, m models.UpdateCategory) (*models.Category, error)
	Delete(ctx context.Context, id string) error
}

type BrandI interface {
	Create(context.Context, models.CreateBrand) (*models.Brand, error)
	GetList(context.Context, models.Pagination) (*models.Pagination, error)
	GetByID(context.Context, string) (*models.Brand, error)
	Update(ctx context.Context, m models.UpdateBrand) (*models.Brand, error)
	Delete(ctx context.Context, id string) error
}
