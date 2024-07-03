package service

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
)

type Product struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
}

func NewProduct(storage storage.StorageInterface, log logs.LoggerInterface) *Product {
	return &Product{storage: storage, log: log}
}

func (srv *Product) Create(ctx context.Context, m models.CreateProduct) status.Status {
	// check price validity
	if m.Price <= 0 {
		return status.StatusBadRequest.AddError("price", status.ErrBadValue)
	}
	var errorListForStatus = make(map[string]status.StatusError)

	// check category existence
	if m.CategoryID != nil {
		if _, err := srv.storage.Category().GetByID(ctx, *m.CategoryID); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errorListForStatus["category_id"] = status.ErrNotFound
			} else {
				return status.StatusInternal
			}
		}
	}

	// check brand existance
	if m.BrandID != nil {
		if _, err := srv.storage.Brand().GetByID(ctx, *m.BrandID); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errorListForStatus["brand_id"] = status.ErrNotFound
			} else {
				return status.StatusInternal
			}
		}
	}

	// check Image existance
	if m.ImageID != nil {
		if _, err := srv.storage.Files().GetByID(ctx, *m.ImageID); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errorListForStatus["image_id"] = status.ErrNotFound
			} else {
				return status.StatusInternal
			}
		}
	}

	// check video existance
	if m.VideoID != nil {
		if _, err := srv.storage.Files().GetByID(ctx, *m.VideoID); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errorListForStatus["video_id"] = status.ErrNotFound
			} else {
				return status.StatusInternal
			}
		}
	}

	// check product files existance
	for _, file := range m.ProductFiles {
		if _, err := srv.storage.Files().GetByID(ctx, file.FileID); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				errorListForStatus[file.FileID] = status.ErrNotFound
			} else {
				return status.StatusInternal
			}
		}
	}

	if len(errorListForStatus) > 0 {
		resp := status.StatusBadRequest
		for k, v := range errorListForStatus {
			resp = resp.AddError(k, v)
		}

		return resp
	}

	product, err := srv.storage.Product().Create(ctx, m)
	if err != nil {
		return status.StatusInternal
	}

	return status.StatusOk.AddData(product)
}

func (srv *Product) GetByID(ctx context.Context, product_id string) status.Status {
	product, err := srv.storage.Product().GetByID(ctx, product_id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusNotFound
		}
		return status.StatusInternal
	}

	return status.StatusOk.AddData(product)
}
