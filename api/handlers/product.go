package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// ProductCreate handles the creation of a new product.
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
// @failure		500 {object}	status.Status	"Internal server error"
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

// ProductGetList handles fetching a list of products with pagination.
// @id			ProductGetList
// @router		/api/v1/product [get]
// @description	Get list of products
// @summary		Get list of products
// @tags		product
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @success		200	{object}	status.Status	"product list"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) ProductGetList(c *gin.Context) {
	var pagination models.Pagination
	if err := c.Bind(&pagination); err != nil {
		v1.ValidateError(c, err, pagination)
		return
	}

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().GetList(ctx, pagination)
	v1.response(c, res)
}

// ProductGetByID handles fetching a product by its ID.
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

// ProductDelete handles deleting a product by its ID.
// @id			ProductDelete
// @router		/api/v1/product/{id} [delete]
// @description	Delete product by id
// @summary		Delete product by id
// @tags		product
// @param		id	path	string	true	"product id"
// @success		200	{object}	status.Status	"product deleted"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) ProductDelete(c *gin.Context) {
	productID := c.Param("id")
	if !helper.IsValidUUID(productID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().Delete(ctx, productID)
	v1.response(c, res)
}

// ProductFilesByID get product files
// @id			ProductFilesByID
// @router		/api/v1/product/{id}/files [get]
// @description	Get product files
// @summary		Get product files
// @tags		product
// @param		id	path	string	true	"product id"
// @success		200	{object}	status.Status	"product files"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) ProductFilesByID(c *gin.Context) {
	productID := c.Param("id")
	if !helper.IsValidUUID(productID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().GetFilesByID(ctx, productID)
	v1.response(c, res)
}

// ProductInstallmentsByID get product installments
// @id			ProductInstallmentsByID
// @router		/api/v1/product/{id}/installments [get]
// @description	Get product installments
// @summary		Get product installments
// @tags		product
// @param		id	path	string	true	"product id"
// @success		200	{object}	status.Status	"product installments"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) ProductInstallmentsByID(c *gin.Context) {
	productID := c.Param("id")
	if !helper.IsValidUUID(productID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Product().GetInstallmentsByID(ctx, productID)
	v1.response(c, res)
}
