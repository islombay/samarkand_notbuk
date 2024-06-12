package service

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
)

type Seller struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
}

func NewSeller(storage storage.StorageInterface, log logs.LoggerInterface) *Seller {
	return &Seller{storage: storage, log: log}
}

func (srv *Seller) Create(ctx context.Context, req models.CreateSeller) status.Status {
	seller, err := srv.storage.Seller().Create(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(seller)
}

func (srv *Seller) GetList(ctx context.Context, p models.Pagination) status.Status {
	pagination, err := srv.storage.Seller().GetList(ctx, p)
	if err != nil {
		return status.StatusInternal
	}

	return status.StatusOk.AddData(pagination.Data).AddCount(pagination.Count)
}

func (srv *Seller) GetByID(ctx context.Context, id string) status.Status {
	seller, err := srv.storage.Seller().GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(seller)
}

func (srv *Seller) Delete(ctx context.Context, id string) status.Status {
	if err := srv.storage.Seller().Delete(ctx, id); err != nil {
		return status.StatusInternal
	}

	return status.StatusOk
}

func (srv *Seller) Update(ctx context.Context, req models.UpdateSeller) status.Status {
	seller, err := srv.storage.Seller().Update(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound.AddError("id", status.ErrNotFound)
		} else if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(seller)
}
