package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
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
