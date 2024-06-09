package service

import (
	"context"
	"errors"

	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/auth"
	"github.com/islombay/noutbuk_seller/pkg/helper"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/storage"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	storage storage.StorageInterface
	log     logs.LoggerInterface
}

func NewAuth(storage storage.StorageInterface, log logs.LoggerInterface) *Auth {
	return &Auth{storage: storage, log: log}
}

func (srv *Auth) LoginAdmin(ctx context.Context, req models.Login) status.Status {

	if !helper.IsValidPhone(req.PhoneNumber) {
		return status.StatusBadPhone
	}

	user, err := srv.storage.User().GetStaffByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusUnauthorized
		}
		return status.StatusInternal
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return status.StatusUnauthorized
		}
		srv.log.Error("could not compare hash and password", logs.Error(err))
		return status.StatusInternal
	}

	tkn := auth.Token{
		UserID: user.ID,
		Role:   auth.RoleAdmin,
	}

	token, err := auth.GenerateToken(tkn)
	if err != nil {
		srv.log.Error("could not generate token", logs.Error(err))
		return status.StatusInternal
	}

	st := status.StatusTokenResponse
	st.Data = token

	return st
}
