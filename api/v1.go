package api

import (
	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/handlers"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/service"
)

func NewV1(
	r *gin.RouterGroup,
	service service.ServiceInterface,
	log logs.LoggerInterface,
) {
	handler := handlers.New(log, service)

	v1 := r.Group("/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login_admin", handler.LoginAdmin)
		auth.GET("/access_token", handler.IsRefreshToken(), handler.GetAccessToken)
	}

	category := v1.Group("/category")
	{
		category.POST("", handler.MiddlewareIsAdmin(), handler.CategoryCreate)
		category.DELETE("", handler.MiddlewareIsAdmin(), handler.CategoryDelete)
		category.PUT("", handler.MiddlewareIsAdmin(), handler.CategoryUpdate)
		category.GET("", handler.CategoryGetList)
		category.GET("/:id", handler.CategoryGetByID)
	}

	brand := v1.Group("/brand")
	{
		brand.POST("", handler.MiddlewareIsAdmin(), handler.BrandCreate)
		brand.DELETE("", handler.MiddlewareIsAdmin(), handler.BrandDelete)
		brand.PUT("", handler.MiddlewareIsAdmin(), handler.BrandUpdate)
		brand.GET("", handler.BrandGetList)
		brand.GET("/:id", handler.BrandGetByID)
	}

	v1.GET("/ping", handler.Ping)
}
