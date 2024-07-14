package postgresql

import (
	"context"
	"errors"
	"time"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type productRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewProductRepo(db *gorm.DB, log logs.LoggerInterface) storage.ProductI {
	return &productRepo{db: db, log: log}
}

func (db *productRepo) Create(ctx context.Context, m models.CreateProduct) (*models.Product, error) {
	product := models.Product{
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,

		CategoryID: m.CategoryID,
		BrandID:    m.BrandID,
		ImageID:    m.ImageID,
		VideoID:    m.VideoID,
	}
	// fmt.Println(product)
	res := db.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			db.log.Error("could not create product", logs.Error(err))
			return err
		}

		for _, file := range m.ProductFiles {
			productFile := models.ProductFile{
				ProductID: product.ID,
				FileID:    file.FileID,
			}
			if err := tx.Create((&productFile)).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return storage.ErrDuplicateProductFile
				} else {
					db.log.Error("could not create product file", logs.Error(err))
					tx.Rollback()
					return err
				}
			}
		}

		for _, installment := range m.ProductInstallments {
			productInstallment := models.ProductInstallment{
				ProductID: product.ID,
				Price:     installment.Price,
				Period:    installment.Period,
			}
			if err := tx.Create(&productInstallment).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return storage.ErrDuplicateProductInstallment
				}
				db.log.Error("could not create product installment", logs.Error(err))
				tx.Rollback()
				return err
			}
		}

		return nil
	})

	return &product, res
}

func (db *productRepo) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product

	m := db.db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})

	if err := m.
		Preload("Image", "deleted_at is null").
		Preload("Video", "deleted_at is null").
		Where("deleted_at IS NULL AND id = ?", id).
		First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find product by id", logs.Error(err))
		return nil, err
	}

	return &product, nil
}

func (db *productRepo) GetList(ctx context.Context, p models.Pagination) (*models.Pagination, error) {
	var products []models.Product

	var tx *gorm.DB = db.db.Where("deleted_at is null")

	if p.Query != "" {
		tx = tx.Where("name ilike ?", "%"+p.Query+"%")
	}

	if p.CategoryID != "" {
		subQuery := db.db.Model(&models.Category{}).
			Select("id").
			Where("parent_id = ? or id = ?", p.CategoryID, p.CategoryID).
			Where("deleted_at is null")

		tx = tx.Where("category_id in (?)", subQuery)
	}

	if p.BrandID != "" {
		tx = tx.Where("brand_id = ?", p.BrandID)
	}

	p.Count = 0

	if res := tx.Model(&models.Product{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of product", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).
		Offset(p.Offset).
		Preload("Image", "deleted_at is null").
		Preload("Video", "deleted_at is null").
		Preload("Category", "deleted_at is null").
		Preload("Brand", "deleted_at is null").
		Order("created_at desc").Find(&products); res.Error != nil {
		db.log.Error("could not get product list", logs.Error(res.Error))
		return nil, res.Error
	}
	p.Data = products

	return &p, nil
}

func (db *productRepo) Delete(ctx context.Context, id string) error {
	err := db.db.Model(&models.Product{ID: id}).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
	if err != nil {
		db.log.Error("could not delete product", logs.Error(err))
		return err
	}

	return nil
}

// GetFilesByID get list of files for specific product
func (db *productRepo) GetFilesByID(ctx context.Context, id string) ([]models.ProductFile, error) {
	var files []models.ProductFile

	if err := db.db.
		WithContext(ctx).
		Joins("join products on products.id = product_files.product_id and products.deleted_at is null").
		Joins("join files on product_files.file_id = files.id and files.deleted_at is null").
		Where("product_files.product_id = ? and product_files.deleted_at is null", id).
		Select("product_files.file_id", "files.file_url", "product_files.created_at", "product_files.updated_at").
		Find(&files).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find files of product", logs.Error(err), logs.String("product_id", id))
		return nil, err
	}
	return files, nil
}

// GetInstallmentsByID get list of installments for specific product
func (db *productRepo) GetInstallmentsByID(ctx context.Context, id string) ([]models.ProductInstallment, error) {
	var installments []models.ProductInstallment

	if err := db.db.
		WithContext(ctx).
		Joins("join products on products.id = product_installments.product_id and products.deleted_at is null").
		Where("product_installments.product_id = ? and product_installments.deleted_at is null", id).
		Select("product_installments.period", "product_installments.price", "product_installments.created_at", "product_installments.updated_at").
		Find(&installments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find installments of product", logs.Error(err), logs.String("product_id", id))
		return nil, err
	}
	return installments, nil
}
