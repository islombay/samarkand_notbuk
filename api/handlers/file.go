package handlers

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/models"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/helper"
	"github.com/islombay/noutbuk_seller/pkg/logs"
)

// UploadFile
// @id			UploadFile
// @router		/api/v1/files/upload [post]
// @description	Upload file
// @summary		Upload file
// @tags 		files
// @security	ApiKeyAuth
// @accept		json
// @product		json
// @param		ile 	formData	models.UploadFile	true 	"Upload file request"
// @param		file	formData	file				true	"File"
// @success		200	{object}	status.Status	"File info"
// @failure		400	{object}	status.Status	"Bad request"
// @failure		500 {object}	status.Status	"Internel server error"
func (v1 *Handler) UploadFile(c *gin.Context) {
	var m models.UploadFile
	if err := c.Bind(&m); err != nil {
		v1.ValidateError(c, err, m)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Files().Create(ctx, m)
	v1.response(c, res)
}

// FileGetByID
// @id			FileGetByID
// @router		/api/v1/files/{id} [get]
// @description	Get file by id
// @summary		Get file by id
// @tags		files
// @param		id	path	string	true	"file id"
// @success		200	{object}	status.Status	"file"
// @failure		400	{object}	status.Status	"Invalid id"
// @failure		404	{object}	status.Status	"Not found"
// @failure 	500 {object}	status.Status	"Internal server error"
func (v1 *Handler) FileGetByID(c *gin.Context) {
	file_id := c.Param("id")
	if !helper.IsValidUUID(file_id) {
		v1.response(c, status.StatusBadID)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Files().GetByID(ctx, file_id)

	fmt.Println(res.FileObjectInfo.ContentType)

	c.Writer.Header().Set("Content-Disposition", "inline; filename="+res.FileObjectInfo.Key)
	c.Writer.Header().Set("Content-Type", res.FileObjectInfo.ContentType)
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", res.FileObjectInfo.Size))

	if _, err := io.Copy(c.Writer, res.FileObject); err != nil {
		v1.log.Error("could not streamline the file to client", logs.Error(err))
		v1.response(c, status.StatusInternal)
	}

	res.FileObject.Close()
}

// FileGetList
// @id			FileGetList
// @router		/api/v1/files [get]
// @description	Get files list
// @summary		Get files list
// @tags		files
// @param		pagination	query	models.Pagination	false 	"Pagination"
// @success		200	{object}	status.Status	"List of files"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) FileGetList(c *gin.Context) {
	var pagination models.Pagination
	c.ShouldBind(&pagination)

	pagination.Fix()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := v1.service.Files().GetList(ctx, pagination)
	v1.response(c, res)
}

// FileDelete
// @id			FileDelete
// @router		/api/v1/files [delete]
// @security 	ApiKeyAuth
// @tags		files
// @description	Delete file
// @summary		Delete file
// @param		id	body	models.DeleteRequest	true	"Delete"
// @success		200 {object}	status.Status	"Success"
// @failure		400	{object}	status.Status	"Bad request / Bad id"
// @failure		500	{object}	status.Status	"Internal server error"
func (v1 *Handler) FileDelete(c *gin.Context) {
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

	res := v1.service.Files().Delete(ctx, m.ID)
	v1.response(c, res)
}
