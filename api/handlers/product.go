package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// ProductCreate
// @id			ProductCreate
// @router		/api/v1/product [post]
// @description	Create product
// @summary		Create product
// @tags 		product
// @security	ApiKeyAuth
// @accept		json
// @product		json
// @param		product 	body	models.CreateProduct	true 	"Create product request"
// @success		200	{object}	status.Status	"Product"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internel server error"
func (v1 *Handler) ProductCreate(c *gin.Context) {
	var m models.CreateProduct
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	for _, file := range m.ProductFiles {
		if !helper.IsValidUUID(file.FileID) {
			v1.response(c, status.StatusBadRequest.AddError(file.FileID, status.ErrUUID))
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().Create(ctx, m)
	v1.response(c, res)
}

// ProductGetByID
// @id			ProductGetByID
// @router		/api/v1/product/{id} [get]
// @description	Get product by id
// @summary		Get product by id
// @tags		product
// @param		id	path	string	true	"product id"
// @success		200	{object}	status.Status	"product"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) ProductGetByID(c *gin.Context) {
	productID := c.Param("id")
	if !helper.IsValidUUID(productID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().GetByID(ctx, productID)
	v1.response(c, res)
}
