package service

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
)

type Category struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
}

func NewCategory(storage storage.StorageInterface, log logs.LoggerInterface) *Category {
	return &Category{storage: storage, log: log}
}

func (srv *Category) Create(ctx context.Context, req models.CreateCategory) status.Status {
	var parent *models.Category
	var err error
	if req.ParentID != nil {
		parent, err = srv.storage.Category().GetByID(ctx, *req.ParentID)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				return status.StatusNotFound.AddError("parent_id", "not_found")
			}
			return status.StatusInternal
		}
	}

	category, err := srv.storage.Category().Create(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	category.Parent = parent

	return status.StatusOk.AddData(category)
}

func (srv *Category) GetList(ctx context.Context, p models.Pagination) status.Status {
	pagination, err := srv.storage.Category().GetList(ctx, p)
	if err != nil {
		return status.StatusInternal
	}

	return status.StatusOk.AddData(pagination.Data).AddCount(pagination.Count)
}

func (srv *Category) GetByID(ctx context.Context, id string) status.Status {
	category, err := srv.storage.Category().GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(category)
}

func (srv *Category) Update(ctx context.Context, req models.UpdateCategory) status.Status {
	var (
		err error
	)
	if req.ParentID != nil {
		_, err = srv.storage.Category().GetByID(ctx, *req.ParentID)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				return status.StatusNotFound.AddError("parent_id", "not_found")
			}
			return status.StatusInternal
		}
	}

	category, err := srv.storage.Category().Update(ctx, req)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound.AddError("id", "not_found")
		} else if errors.Is(err, storage.ErrAlreadyExists) {
			return status.StatusAlreadyExists
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(category)
}

func (srv *Category) Delete(ctx context.Context, id string) status.Status {
	if err := srv.storage.Category().Delete(ctx, id); err != nil {
		return status.StatusInternal
	}

	return status.StatusOk
}
