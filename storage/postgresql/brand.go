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

type brandRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewBrandRepo(db *gorm.DB, log logs.LoggerInterface) storage.BrandI {
	return &brandRepo{db: db, log: log}
}

func (db *brandRepo) Create(ctx context.Context, m models.CreateBrand) (*models.Brand, error) {
	brand := models.Brand{
		Name: m.Name,
	}

	result := db.db.Select("name").Create(&brand)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not create brand", logs.Error(result.Error))
		return nil, result.Error
	}

	return &brand, nil
}

func (db *brandRepo) GetList(ctx context.Context, p models.Pagination) (*models.Pagination, error) {
	var brands []models.Brand

	var tx *gorm.DB = db.db.Where("deleted_at is null")

	if p.Query != "" {
		tx = tx.Where("name ilike ?", "%"+p.Query+"%")
	}

	p.Count = 0

	if res := tx.Model(&models.Brand{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of brand", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).Offset(p.Offset).Order("created_at desc").Find(&brands); res.Error != nil {
		db.log.Error("could not get brand list", logs.Error(res.Error))
		return nil, res.Error
	}
	p.Data = brands

	return &p, nil
}

func (db *brandRepo) GetByID(ctx context.Context, id string) (*models.Brand, error) {
	var brand models.Brand

	if res := db.db.Where("deleted_at is null").First(&brand, "id = ?", id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find brand by id", logs.Error(res.Error))
		return nil, res.Error
	}

	return &brand, nil
}

func (db *brandRepo) Delete(ctx context.Context, id string) error {
	err := db.db.Model(&models.Brand{ID: id}).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
	if err != nil {
		db.log.Error("could not delete brand", logs.Error(err))
		return err
	}

	return nil
}

func (db *brandRepo) Update(ctx context.Context, m models.UpdateBrand) (*models.Brand, error) {
	if err := db.db.Model(&models.Brand{ID: m.ID}).Where("deleted_at is null").Updates(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not update brand", logs.Error(err))
		return nil, err
	}

	return db.GetByID(ctx, m.ID)
}
