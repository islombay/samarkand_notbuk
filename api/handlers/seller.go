package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// SellerCreate
// @id			SellerCreate
// @router		/api/v1/seller [post]
// @description	Create seller
// @summary		Create seller
// @tags 		seller
// @security	ApiKeyAuth
// @accept		json
// @product		json
// @param		seller 	body	models.CreateSeller	true 	"Create seller request"
// @success		200	{object}	status.Status	"seller"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internel server error"
func (v1 *Handler) SellerCreate(c *gin.Context) {
	var m models.CreateSeller
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Seller().Create(ctx, m)
	v1.response(c, res)
}

// SellerGetList
// @id			SellerGetList
// @router		/api/v1/seller [get]
// @description	Get seller list
// @summary		Get seller list
// @tags		seller
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @success		200	{object}	status.Status	"List of sellers"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) SellerGetList(c *gin.Context) {
	var pagination models.Pagination
	c.ShouldBind(&pagination)

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Seller().GetList(ctx, pagination)
	v1.response(c, res)
}

// SellerGetByID
// @id			SellerGetByID
// @router		/api/v1/seller/{id} [get]
// @description	Get seller by id
// @summary		Get seller by id
// @tags		seller
// @param		id	path	string	true	"seller id"
// @success		200	{object}	status.Status	"seller"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) SellerGetByID(c *gin.Context) {
	sellerID := c.Param("id")
	if !helper.IsValidUUID(sellerID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Seller().GetByID(ctx, sellerID)
	v1.response(c, res)
}

// SellerUpdate
// @id			BrandUpdate
// @router		/api/v1/seller [put]
// @description Update seller. Update only given values
// @summary		Update seller
// @tags		seller
// @security	ApiKeyAuth
// @param		seller	body	models.UpdateSeller	true	"Update"
// @success		200	{object}	status.Status	"seller"
// @failure		400 {object}	status.Status	"Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internal server error"
func (v1 *Handler) SellerUpdate(c *gin.Context) {
	var m models.UpdateSeller
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.ID) {
		v1.response(c, status.StatusBadID.AddError("id", status.ErrInvalid))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Seller().Update(ctx, m)
	v1.response(c, res)
}

// SellerDelete
// @id			SellerDelete
// @router		/api/v1/seller [delete]
// @security 	ApiKeyAuth
// @tags		seller
// @description	Delete seller
// @summary		Delete seller
// @param		id	body	models.DeleteRequest	true	"Delete"
// @success		200 {object}	status.Status	"Success"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) SellerDelete(c *gin.Context) {
	var m models.DeleteRequest
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.ID) {
		v1.response(c, status.StatusBadID.AddError("id", status.ErrInvalid))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Seller().Delete(ctx, m.ID)
	v1.response(c, res)
}
