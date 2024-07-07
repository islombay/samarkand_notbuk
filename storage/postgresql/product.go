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

	m := db.db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})

	if err := m.
		Preload("Category").
		Preload("Brand").
		Preload("Image").
		Preload("Video").
		Preload("ProductFiles").
		Preload("ProductInstallments").
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

	p.Count = 0

	if res := tx.Model(&models.Product{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of product", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).Offset(p.Offset).Preload("Image").Preload("Video").Order("created_at desc").Find(&products); res.Error != nil {
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
