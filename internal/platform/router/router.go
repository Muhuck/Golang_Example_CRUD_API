package router

import (
	"crud-api/internal/category"

	"github.com/gin-gonic/gin"
)

func Setup(categoryHandler *category.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	categoryHandler.Register(api.Group("/categories"))
	return r
}
