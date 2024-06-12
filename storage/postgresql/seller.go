package postgresql

import (
	"context"
	"errors"
	"time"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"gorm.io/gorm"
)

type sellerRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewSellerRepo(db *gorm.DB, log logs.LoggerInterface) storage.SellerI {
	return &sellerRepo{db: db, log: log}
}

func (db *sellerRepo) Create(ctx context.Context, m models.CreateSeller) (*models.Seller, error) {
	seller := models.Seller{
		FirstName:   m.FirstName,
		LastName:    m.LastName,
		PhoneNumber: m.PhoneNumber,
	}

	result := db.db.Select("first_name", "last_name", "phone_number").Create(&seller)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not create seller", logs.Error(result.Error))
		return nil, result.Error
	}

	return &seller, nil
}

func (db *sellerRepo) GetList(ctx context.Context, p models.Pagination) (*models.Pagination, error) {
	var sellers []models.Seller

	var tx *gorm.DB = db.db.Where("deleted_at is null")

	if p.Query != "" {
		tx = tx.Where("first_name ilike ?", "%"+p.Query+"%").Where(
			"last_name ilike ?", "%"+p.Query+"%",
		).Where("phone_number ilike ?", "%"+p.Query+"%")
	}

	p.Count = 0

	if res := tx.Model(&models.Seller{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of seller", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).Offset(p.Offset).Order("created_at desc").Find(&sellers); res.Error != nil {
		db.log.Error("could not get seller list", logs.Error(res.Error))
		return nil, res.Error
	}
	p.Data = sellers

	return &p, nil
}

func (db *sellerRepo) GetByID(ctx context.Context, id string) (*models.Seller, error) {
	var seller models.Seller

	if res := db.db.Where("deleted_at is null").First(&seller, "id = ?", id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find seller by id", logs.Error(res.Error))
		return nil, res.Error
	}

	return &seller, nil
}

func (db *sellerRepo) Delete(ctx context.Context, id string) error {
	err := db.db.Model(&models.Seller{ID: id}).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
	if err != nil {
		db.log.Error("could not delete seller", logs.Error(err))
		return err
	}

	return nil
}

func (db *sellerRepo) Update(ctx context.Context, m models.UpdateSeller) (*models.Seller, error) {
	if err := db.db.Model(&models.Seller{ID: m.ID}).Where("deleted_at is null").Updates(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not update seller", logs.Error(err))
		return nil, err
	}

	return db.GetByID(ctx, m.ID)
}
