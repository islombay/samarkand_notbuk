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
		auth.POST("/login", handler.Login)
		auth.POST("/verify", handler.VerifyLogin)
		auth.POST("/profile", handler.LoginProfile)

		auth.GET("/access_token", handler.IsRefreshToken(), handler.GetAccessToken)
	}

	category := v1.Group("/category")
	{
		category.POST("", handler.MiddlewareIsAdmin(), handler.CategoryCreate)
		category.DELETE("", handler.MiddlewareIsAdmin(), handler.CategoryDelete)
		category.PUT("", handler.MiddlewareIsAdmin(), handler.CategoryUpdate)
		category.GET("", handler.CategoryGetList)
		category.GET("/:id", handler.CategoryGetByID)

		v1.GET("/subcategory", handler.CategoryGetListSub)
		v1.PUT("/subcategory", handler.MiddlewareIsAdmin(), handler.CategoryUpdate)
		v1.DELETE("/subcategory", handler.MiddlewareIsAdmin(), handler.CategoryDelete)
	}

	brand := v1.Group("/brand")
	{
		brand.POST("", handler.MiddlewareIsAdmin(), handler.BrandCreate)
		brand.DELETE("", handler.MiddlewareIsAdmin(), handler.BrandDelete)
		brand.PUT("", handler.MiddlewareIsAdmin(), handler.BrandUpdate)
		brand.GET("", handler.BrandGetList)
		brand.GET("/:id", handler.BrandGetByID)
	}

	seller := v1.Group("/seller")
	{
		seller.POST("", handler.MiddlewareIsAdmin(), handler.SellerCreate)
		seller.DELETE("", handler.MiddlewareIsAdmin(), handler.SellerDelete)
		seller.PUT("", handler.MiddlewareIsAdmin(), handler.SellerUpdate)
		seller.GET("", handler.SellerGetList)
		seller.GET("/:id", handler.SellerGetByID)
	}

	files := v1.Group("/files")
	{
		files.POST("/upload", handler.MiddlewareIsAdmin(), handler.UploadFile)
		files.GET("/:id", handler.FileGetByID)
		files.GET("", handler.FileGetList)
		files.DELETE("", handler.MiddlewareIsAdmin(), handler.FileDelete)
	}

	product := v1.Group("/product")
	{
		product.POST("", handler.MiddlewareIsAdmin(), handler.ProductCreate)
		product.DELETE("", handler.MiddlewareIsAdmin(), handler.ProductDelete)
		product.GET("", handler.ProductGetList)
		product.GET("/:id", handler.ProductGetByID)
		product.GET("/:id/files", handler.ProductFilesByID)
		product.GET("/:id/installments", handler.ProductInstallmentsByID)
	}

	v1.GET("/ping", handler.Ping)
}
