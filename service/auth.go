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

	access_token, err := auth.GenerateToken(tkn)
	if err != nil {
		srv.log.Error("could not generate access token", logs.Error(err))
		return status.StatusInternal
	}

	refresh_token, err := auth.GenerateTokenRefresh(tkn)
	if err != nil {
		srv.log.Error("could not generate refresh token", logs.Error(err))
		return status.StatusInternal
	}

	st := status.StatusTokenResponse
	st.Data = map[string]string{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	}

	return st
}

func (srv *Auth) Login(ctx context.Context, req models.Login) status.Status {
	if !helper.IsValidPhone(req.PhoneNumber) {
		return status.StatusBadPhone
	}

	// /auth/login
	// {
	//	"phone_number": ""
	// "header device_id": ""
	// }
	//	create user with phone number if not exists
	//	if redis (device_id and user_id) exists:
	//		return request_id
	//	else:
	//		send sms code to phone number
	//		set redis (request_id user_id code, need_phone:true device_id)
	//		return request_id
	//	GOTO: /verify

	// /auth/verify
	// {
	//	"request_id",
	//	"code"
	// }
	//	get redis (by request_id)
	//	if user.code == redis.code:
	//		if user.name not set:
	//			return code:201 request_id
	//		if user.name set:
	//			return access_token refresh_token code:200
	//	else:
	//		return bad code code:406
	//	}

	// /auth/profile
	// {
	// "request_id",
	//	"first_name",
	//	"last_name"
	// }
	//	get redis (by request_id)
	//	if redis.

	_, err := srv.storage.User().GetClientByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return status.StatusUnauthorized
		}
		return status.StatusInternal
	}

	return status.Status{}
}

func (srv *Auth) GetAccessToken(ctx context.Context, req auth.Token) status.Status {
	token := auth.Token{}

	if req.Role == auth.RoleAdmin {
		user, err := srv.storage.User().GetStaffByID(ctx, req.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				return status.StatusUnauthorized
			}
			return status.StatusInternal
		}

		token.UserID = user.ID
		token.Role = req.Role
	}

	access_token, err := auth.GenerateToken(token)
	if err != nil {
		srv.log.Error("could not generate access token", logs.Error(err))
		return status.StatusInternal
	}

	return status.StatusTokenResponse.AddData(
		map[string]string{
			"access_token": access_token,
		},
	)
}
