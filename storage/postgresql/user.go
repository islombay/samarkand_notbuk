package postgresql

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"gorm.io/gorm"
)

type UserRepo struct {
	db  *gorm.DB
	log logs.LoggerInterface
}

func NewUserRepo(db *gorm.DB, log logs.LoggerInterface) *UserRepo {
	return &UserRepo{db: db, log: log}
}

func (db *UserRepo) GetStaffByPhoneNumber(ctx context.Context, phone_number string) (*models.Staff, error) {
	var user models.Staff
	tx := db.db.Where(&models.Staff{PhoneNumber: phone_number}).Where("deleted_at is null").First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not get staff by phone number", logs.Error(tx.Error))
		return nil, tx.Error
	}

	return &user, nil
}

func (db *UserRepo) GetClientByPhoneNumber(ctx context.Context, phone_number string) (*models.Client, error) {
	var user models.Client
	tx := db.db.Model(&models.Client{PhoneNumber: phone_number}).Where("deleted_at is null").First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not get client by phone number", logs.Error(tx.Error))
		return nil, tx.Error
	}

	return &user, nil
}

func (db *UserRepo) GetStaffByID(ctx context.Context, id string) (*models.Staff, error) {
	var user models.Staff
	tx := db.db.Model(&models.Staff{ID: id}).Where("deleted_at is null").First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, storage.ErrNotFound
		}
		db.log.Error("could not get staff by id", logs.Error(tx.Error))
		return nil, tx.Error
	}

	return &user, nil
}
