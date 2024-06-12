package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/auth"
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
