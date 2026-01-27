package main

import (
	"crud-api/config"
	"crud-api/internal/category"
	"crud-api/internal/platform/database"
	"crud-api/internal/platform/router"
	"crud-api/internal/product"
	"os"
)

func main() {
	config.Load()

	db := database.Connect(os.Getenv("DATABASE_URL"))
	db.AutoMigrate(
		&category.Category{},
		&product.Product{},
	)

	repo_category := category.NewRepository(db)
	service_category := category.NewService(repo_category)
	handler_category := category.NewHandler(service_category)

	repo_product := product.NewRepository(db)
	service_product := product.NewService(repo_product)
	handler_product := product.NewHandler(service_product)

	r := router.Setup(handler_category, handler_product)
	r.Run(":" + os.Getenv("PORT"))
}
