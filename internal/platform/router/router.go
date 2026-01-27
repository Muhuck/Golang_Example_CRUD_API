package router

import (
	"crud-api/internal/category"
	"crud-api/internal/product"

	"github.com/gin-gonic/gin"
)

func Setup(categoryHandler *category.Handler, productHandler *product.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	categoryHandler.Register(api.Group("/categories"))
	productHandler.Register(api.Group("/products"))
	return r
}
