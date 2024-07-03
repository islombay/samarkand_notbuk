package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/islombay/noutbuk_seller/api/docs"
	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/helper"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(
	r *gin.Engine,
	service service.ServiceInterface,
	cfg config.Config,
	log logs.LoggerInterface,
) {
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	docs.SwaggerInfo.Title = "Noutbuk"
	docs.SwaggerInfo.Description = "Noutbuk"
	docs.SwaggerInfo.Version = "1.0"

	r.Use(customCORSMiddleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uuid", helper.UUIDValidator)
	}

	api := r.Group("/api")

	NewV1(api, service, log)

	docs.SwaggerInfo.Host = cfg.Server.Public
	// if cfg.ENV == config.LocalMode {
	// 	docs.SwaggerInfo.Host = cfg.Server.Public
	// } else if cfg.ENV == config.ProdMode {
	// 	docs.SwaggerInfo.Host = cfg.Server.Public
	// }

	r.GET("/sw/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
	))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
