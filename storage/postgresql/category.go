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

type categoryRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewCategoryRepo(db *gorm.DB, log logs.LoggerInterface) storage.CategoryI {
	return &categoryRepo{db: db, log: log}
}

func (db *categoryRepo) Create(ctx context.Context, m models.CreateCategory) (*models.Category, error) {
	cat := models.Category{
		NameUz:   m.NameUz,
		ParentID: m.ParentID,
	}

	result := db.db.Select("name_uz", "parent_id").Create(&cat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not create category", logs.Error(result.Error))
		return nil, result.Error
	}

	return &cat, nil
}

func (db *categoryRepo) GetList(ctx context.Context, p models.Pagination, onlySub bool) (*models.Pagination, error) {
	var categories []models.Category

	var tx *gorm.DB = db.db.Where("deleted_at is null")

	if onlySub {
		tx = tx.Where("parent_id is not null")
	} else {
		tx = tx.Where("parent_id is null")
	}

	if p.Query != "" {
		tx = tx.Where("name_uz ilike ?", "%"+p.Query+"%")
	}

	p.Count = 0

	if res := tx.Model(&models.Category{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of category", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).Offset(p.Offset).Order("created_at desc").Find(&categories); res.Error != nil {
		db.log.Error("could not get category list", logs.Error(res.Error))
		return nil, res.Error
	}
	p.Data = categories

	return &p, nil
}

func (db *categoryRepo) GetByID(ctx context.Context, id string) (*models.Category, error) {
	var category models.Category

	if res := db.db.Preload("Parent").Preload("Subcategories", "deleted_at is null").Where("deleted_at is null").First(&category, "id = ?", id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find category by id", logs.Error(res.Error))
		return nil, res.Error
	}

	return &category, nil
}

func (db *categoryRepo) Update(ctx context.Context, m models.UpdateCategory) (*models.Category, error) {
	if err := db.db.Model(&models.Category{ID: m.ID}).Where("deleted_at is null").Updates(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not update category", logs.Error(err))
		return nil, err
	}

	return db.GetByID(ctx, m.ID)
}

func (db *categoryRepo) Delete(ctx context.Context, id string) error {
	err := db.db.Model(&models.Category{ID: id}).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
	if err != nil {
		db.log.Error("could not delete category", logs.Error(err))
		return err
	}

	return nil
}
