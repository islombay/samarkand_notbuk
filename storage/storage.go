package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/islombay/noutbuk_seller/api/models"
)

var (
	ErrNotFound      = fmt.Errorf("err_not_found")
	ErrAlreadyExists = fmt.Errorf("already_exists")

	// Product
	ErrDuplicateProductFile        = errors.New("duplicate_product_file")
	ErrDuplicateProductInstallment = errors.New("duplicate_product_installment")
)

type StorageInterface interface {
	Close()
	User() UserI
	Category() CategoryI
	Brand() BrandI
	Seller() SellerI
	Files() FilesI
	Product() ProductI
}

type UserI interface {
	GetStaffByPhoneNumber(context.Context, string) (*models.Staff, error)
	GetStaffByID(context.Context, string) (*models.Staff, error)

	GetClientByPhoneNumber(context.Context, string) (*models.Client, error)
	CreateClient(ctx context.Context, client models.Client) (*models.Client, error)
	GetClientByID(context.Context, string) (*models.Client, error)
	UpdateClient(ctx context.Context, m models.UpdateClient) (*models.Client, error)
}

type CategoryI interface {
	Create(context.Context, models.CreateCategory) (*models.Category, error)
	GetList(context.Context, models.Pagination, bool) (*models.Pagination, error)
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

type SellerI interface {
	Create(context.Context, models.CreateSeller) (*models.Seller, error)
	GetList(context.Context, models.Pagination) (*models.Pagination, error)
	GetByID(context.Context, string) (*models.Seller, error)
	Update(ctx context.Context, m models.UpdateSeller) (*models.Seller, error)
	Delete(ctx context.Context, id string) error
}

type FilesI interface {
	Create(ctx context.Context, m models.UploadFile) (*models.File, error)
	GetList(context.Context, models.Pagination) (*models.Pagination, error)
	GetByID(context.Context, string) (*models.File, error)
	Delete(ctx context.Context, id string) error
}

type ProductI interface {
	Create(ctx context.Context, m models.CreateProduct) (*models.Product, error)
	GetList(context.Context, models.Pagination) (*models.Pagination, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Delete(ctx context.Context, id string) error

	GetFilesByID(ctx context.Context, id string) ([]models.ProductFile, error)
	GetInstallmentsByID(ctx context.Context, id string) ([]models.ProductInstallment, error)
}
