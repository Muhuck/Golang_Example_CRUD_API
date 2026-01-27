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
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&product.Product{})

	repo := category.NewRepository(db)
	service := category.NewService(repo)
	handler := category.NewHandler(service)

	repo_product := product.NewRepository(db)
	service_product := product.NewService(repo_product)
	handler_product := product.NewHandler(service_product)

	r := router.Setup(handler, handler_product)
	r.Run(":" + os.Getenv("PORT"))
}
