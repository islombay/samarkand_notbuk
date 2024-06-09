package service

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
)

type Brand struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
}

func NewBrand(storage storage.StorageInterface, log logs.LoggerInterface) *Brand {
	return &Brand{storage: storage, log: log}
}

func (srv *Brand) Create(ctx context.Context, req models.CreateBrand) status.Status {
	brand, err := srv.storage.Brand().Create(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(brand)
}

func (srv *Brand) GetList(ctx context.Context, p models.Pagination) status.Status {
	pagination, err := srv.storage.Brand().GetList(ctx, p)
	if err != nil {
		return status.StatusInternal
	}

	return status.StatusOk.AddData(pagination.Data).AddCount(pagination.Count)
}

func (srv *Brand) GetByID(ctx context.Context, id string) status.Status {
	brand, err := srv.storage.Brand().GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(brand)
}

func (srv *Brand) Update(ctx context.Context, req models.UpdateBrand) status.Status {
	brand, err := srv.storage.Brand().Update(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound.AddError("id", "not_found")
		} else if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(brand)
}

func (srv *Brand) Delete(ctx context.Context, id string) status.Status {
	if err := srv.storage.Brand().Delete(ctx, id); err != nil {
		return status.StatusInternal
	}

	return status.StatusOk
}
