package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
)

// CreateCategory
// @id			CreateCategory
// @router		/api/v1/category [post]
// @description	Create category
// @summary		Create category
// @tags 		category
// @security	ApiKeyAuth
// @accept		json
// @product		json
// @param		category 	body	models.CreateCategory	true 	"Create category request"
// @success		200	{object}	status.Status	"Category"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internel server error"
func (v1 *Handler) CategoryCreate(c *gin.Context) {
	var m models.CreateCategory
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if m.ParentID != nil {
		if *m.ParentID == "" {
			m.ParentID = nil
		} else {
			if !helper.IsValidUUID(*m.ParentID) {
				v1.response(c, status.StatusBadID.AddError("parent_id", status.ErrInvalid))
				return
			}
		}
	}

	res := v1.service.Category().Create(ctx, m)
	v1.response(c, res)
}

// CategoryGetList
// @id			CategoryGetList
// @router		/api/v1/category [get]
// @description	Get category list
// @summary		Get category list
// @tags		category
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @param		only_sub	query	bool				false	"Return only sub category ?"
// @success		200	{object}	status.Status	"List of categories"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) CategoryGetList(c *gin.Context) {
	var pagination struct {
		models.Pagination
		OnlySub bool `form:"only_sub"`
	}
	c.ShouldBind(&pagination)

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Category().GetList(ctx, pagination.Pagination, pagination.OnlySub)
	v1.response(c, res)
}

// CategoryGetListSub
// @id			CategoryGetListSub
// @router		/api/v1/subcategory [get]
// @description	Get subcategory list
// @summary		Get subcategory list
// @tags		category
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @param		only_sub	query	bool				false	"Return only sub category ?"
// @success		200	{object}	status.Status	"List of categories"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) CategoryGetListSub(c *gin.Context) {
	var pagination struct {
		models.Pagination
		OnlySub bool `form:"only_sub"`
	}
	c.ShouldBind(&pagination)

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Category().GetList(ctx, pagination.Pagination, true)
	v1.response(c, res)
}

// CategoryGetByID
// @id			CategoryGetByID
// @router		/api/v1/category/{id} [get]
// @description	Get category by id
// @summary		Get category by id
// @tags		category
// @param		id	path	string	true	"category id"
// @success		200	{object}	status.Status	"Category"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) CategoryGetByID(c *gin.Context) {
	categoryID := c.Param("id")
	if !helper.IsValidUUID(categoryID) {
		v1.response(c, status.StatusBadID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Category().GetByID(ctx, categoryID)
	v1.response(c, res)
}

// CategoryUpdate
// @id			CategoryUpdate
// @router		/api/v1/category [put]
// @description Update category. Update only given values
// @summary		Update category
// @tags		category
// @security	ApiKeyAuth
// @param		category	body	models.UpdateCategory	true	"Update"
// @success		200	{object}	status.Status	"Category"
// @failure		400 {object}	status.Status	"Bad id"
// @failure		404 {object}	status.Status	"Not found"
// @failure		500 {object}	status.Status	"Internal server error"
func (v1 *Handler) CategoryUpdate(c *gin.Context) {
	var m models.UpdateCategory
	if err := c.BindJSON(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	if !helper.IsValidUUID(m.ID) {
		v1.response(c, status.StatusBadID.AddError("id", status.ErrInvalid))
		return
	}

	if m.ParentID != nil {
		if *m.ParentID == "" {
			m.ParentID = nil
		} else {
			if !helper.IsValidUUID(*m.ParentID) {
				v1.response(c, status.StatusBadID.AddError("parent_id", status.ErrInvalid))
				return
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Category().Update(ctx, m)
	v1.response(c, res)
}

// CategoryDelete
// @id			CategoryDelete
// @router		/api/v1/category [delete]
// @security 	ApiKeyAuth
// @tags		category
// @description	Delete category
// @summary		Delete category
// @param		id	body	models.DeleteRequest	true	"Delete"
// @success		200 {object}	status.Status	"Success"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) CategoryDelete(c *gin.Context) {
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

	res := v1.service.Category().Delete(ctx, m.ID)
	v1.response(c, res)
}
