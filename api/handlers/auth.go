package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/auth"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// LoginAdmin
// @id 				LoginAdmin
// @router			/api/v1/auth/login_admin [post]
// @summary			Login as admin
// @description 	Login as admin
// @tags			auth
// @accept			json
// @produce			json
// @param			login body models.Login		true "Login request"
func (v1 *Handler) LoginAdmin(c *gin.Context) {
	var m models.Login
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Auth().LoginAdmin(ctx, m)
	v1.response(c, res)
}

// Login
// @id 				Login
// @router			/api/v1/auth/login [post]
// @summary			Login as client
// @description 	Login as client
// @tags			auth
// @accept			json
// @produce			json
// @param			login body models.LoginClient		true "Login request"
func (v1 *Handler) Login(c *gin.Context) {
	var m models.LoginClient
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Auth().Login(ctx, m)
	v1.response(c, res)
}

// VerifyLogin
// @id 				VerifyLogin
// @router			/api/v1/auth/verify [post]
// @summary			Verify login with sms code
// @description 	Verify login with sms code
// @tags			auth
// @accept			json
// @produce			json
// @param			login body models.VerifyLogin		true "Verify request"
func (v1 *Handler) VerifyLogin(c *gin.Context) {
	var m models.VerifyLogin
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.RequestID) {
		v1.response(c, status.StatusBadID.AddError("request_id", status.ErrInvalid))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Auth().Verify(ctx, m)
	v1.response(c, res)
}

// LoginProfile
// @id 				LoginProfile
// @router			/api/v1/auth/profile [post]
// @summary			Add first and last name to profile while authorization
// @description 	Add first and last name to profile while authorization
// @tags			auth
// @accept			json
// @produce			json
// @param			login body models.ProfileLogin		true "Profile request"
func (v1 *Handler) LoginProfile(c *gin.Context) {
	var m models.ProfileLogin
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.RequestID) {
		v1.response(c, status.StatusBadID.AddError("request_id", status.ErrInvalid))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Auth().Profile(ctx, m)
	v1.response(c, res)
}

// GetAccessToken
// @id 				GetAccessToken
// @router			/api/v1/auth/access_token [get]
// @summary			Get new access token
// @description 	Get new access token
// @tags			auth
// @security		ApiKeyAuth
// @accept			json
// @produce			json
func (v1 *Handler) GetAccessToken(c *gin.Context) {
	token, exists := GetContextValue[*auth.Token](c, "token")
	if !exists {
		v1.response(c, status.StatusInternal)
		v1.log.Error("token does not exist in context")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Auth().GetAccessToken(ctx, *token)
	v1.response(c, res)
}
