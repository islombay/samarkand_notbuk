package handlers

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/service"
)

type Handler struct {
	log     logs.LoggerInterface
	service service.ServiceInterface
}

func New(
	log logs.LoggerInterface,
	service service.ServiceInterface,
) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (v1 *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func (v1 *Handler) response(c *gin.Context, data status.Status) {
	c.JSON(data.Code, data)
}

func (v1 *Handler) ValidateError(c *gin.Context, err error, class interface{}) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {

		out := make(map[string]string)
		for _, fe := range ve {
			jsonFieldName := getJSONFieldName(class, fe.Field())
			if jsonFieldName == "" {
				jsonFieldName = fe.Field()
			}
			out[jsonFieldName] = fe.Tag()
		}

		v := status.StatusBadRequest
		v.Error = out
		v1.response(c, v)
		return
	}
	v1.response(c, status.StatusInternal)
}

// getJSONFieldName maps struct field names to their corresponding JSON tag values.
func getJSONFieldName(v interface{}, field string) string {
	r := reflect.TypeOf(v)
	for i := 0; i < r.NumField(); i++ {
		structField := r.Field(i)
		if structField.Name == field {
			return structField.Tag.Get("json")
		}
	}
	return ""
}
