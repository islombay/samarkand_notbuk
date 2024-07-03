package postgresql

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"gorm.io/gorm"
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
				db.log.Error("could not create product file", logs.Error(err))
				tx.Rollback()
				return err
			}
		}

		for _, installment := range m.ProductInstallments {
			productInstallment := models.ProductInstallment{
				ProductID: product.ID,
				Price:     installment.Price,
				Period:    installment.Period,
			}
			if err := tx.Create(&productInstallment).Error; err != nil {
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

	if err := db.db.Preload("Category").Preload("Brand").Preload("Image").Preload("Video").Preload("ProductFiles").Preload("ProductInstallments").Where("deleted_at is null").First(&product, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find product by id", logs.Error(err))
		return nil, err
	}

	return &product, nil
}
