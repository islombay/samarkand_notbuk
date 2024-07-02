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

type filesRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewFilesRepo(db *gorm.DB, log logs.LoggerInterface) storage.FilesI {
	return &filesRepo{db: db, log: log}
}

func (db *filesRepo) Create(ctx context.Context, m models.UploadFile) (*models.File, error) {
	file := models.File{
		ID:      m.ID,
		FileURL: m.FileURL,
	}
	result := db.db.Create(&file)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, storage.ErrAlreadyExists
		}
		db.log.Error("could not create file in db", logs.Error(result.Error))
		return nil, result.Error
	}
	return &file, nil
}

func (db *filesRepo) GetList(ctx context.Context, p models.Pagination) (*models.Pagination, error) {
	var files []models.File

	var tx *gorm.DB = db.db.Where("deleted_at is null")

	p.Count = 0

	if res := tx.Model(&models.File{}).Count(&p.Count); res.Error != nil {
		db.log.Error("could not get the count of files", logs.Error(res.Error))
	}

	if res := tx.Limit(p.Limit).Offset(p.Offset).Order("created_at desc").Find(&files); res.Error != nil {
		db.log.Error("could not get files list", logs.Error(res.Error))
		return nil, res.Error
	}
	p.Data = files

	return &p, nil
}

func (db *filesRepo) GetByID(ctx context.Context, id string) (*models.File, error) {
	var file models.File

	if res := db.db.Where("deleted_at is null").First(&file, "id = ?", id); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not find file by id", logs.Error(res.Error))
		return nil, res.Error
	}

	return &file, nil
}

func (db *filesRepo) Delete(ctx context.Context, id string) error {
	err := db.db.Model(&models.File{ID: id}).Where("deleted_at is null").Update("deleted_at", time.Now()).Error
	if err != nil {
		db.log.Error("could not delete file", logs.Error(err))
		return err
	}

	return nil
}
