package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// BrandCreate
// @id			BrandCreate
// @router		/api/v1/brand [post]
// @description	Create brand
// @summary		Create brand
// @tags 		brand
// @security	ApiKeyAuth
// @accept		json
// @product		json
// @param		brand 	body	models.CreateBrand	true 	"Create brand request"
// @success		200	{object}	status.Status	"Brand"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internel server error"
func (v1 *Handler) BrandCreate(c *gin.Context) {
	var m models.CreateBrand
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Brand().Create(ctx, m)
	v1.response(c, res)
}

// BrandGetList
// @id			BrandGetList
// @router		/api/v1/brand [get]
// @description	Get brand list
// @summary		Get brand list
// @tags		brand
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @success		200	{object}	status.Status	"List of brands"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) BrandGetList(c *gin.Context) {
	var pagination models.Pagination
	c.ShouldBind(&pagination)

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Brand().GetList(ctx, pagination)
	v1.response(c, res)
}

// BrandGetByID
// @id			BrandGetByID
// @router		/api/v1/brand/{id} [get]
// @description	Get brand by id
// @summary		Get brand by id
// @tags		brand
// @param		id	path	string	true	"brand id"
// @success		200	{object}	status.Status	"brand"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) BrandGetByID(c *gin.Context) {
	brandID := c.Param("id")
	if !helper.IsValidUUID(brandID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Brand().GetByID(ctx, brandID)
	v1.response(c, res)
}

// BrandUpdate
// @id			BrandUpdate
// @router		/api/v1/brand [put]
// @description Update brand. Update only given values
// @summary		Update brand
// @tags		brand
// @security	ApiKeyAuth
// @param		brand	body	models.UpdateBrand	true	"Update"
// @success		200	{object}	status.Status	"brand"
// @failure		400 {object}	status.Status	"Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internal server error"
func (v1 *Handler) BrandUpdate(c *gin.Context) {
	var m models.UpdateBrand
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.ID) {
		v1.response(c, status.StatusBadID.AddError("id", "invalid"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Brand().Update(ctx, m)
	v1.response(c, res)
}

// BrandDelete
// @id			BrandDelete
// @router		/api/v1/brand [delete]
// @security 	ApiKeyAuth
// @tags		brand
// @description	Delete brand
// @summary		Delete brand
// @param		id	body	models.DeleteRequest	true	"Delete"
// @success		200 {object}	status.Status	"Success"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) BrandDelete(c *gin.Context) {
	var m models.DeleteRequest
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.ID) {
		v1.response(c, status.StatusBadID.AddError("id", "invalid"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Brand().Delete(ctx, m.ID)
	v1.response(c, res)
}
