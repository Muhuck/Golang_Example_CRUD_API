package router

import (
	"crud-api/internal/category"
	"crud-api/internal/product"
	"crud-api/internal/report"
	"crud-api/internal/transaction"

	"github.com/gin-gonic/gin"
)

func Setup(categoryHandler *category.Handler, productHandler *product.Handler, transactionHandler *transaction.Handler, reportHandler *report.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	categoryHandler.Register(api.Group("/categories"))
	productHandler.Register(api.Group("/products"))
	transactionHandler.Register(api.Group("/transactions"))
	reportHandler.Register(api.Group("/report"))
	return r
}
