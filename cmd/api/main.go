package main

import (
	"crud-api/config"
	"crud-api/internal/category"
	"crud-api/internal/platform/database"
	"crud-api/internal/platform/router"
	"os"
)

func main() {
	config.Load()

	db := database.Connect(os.Getenv("DATABASE_URL"))
	db.AutoMigrate(&category.Category{})

	repo := category.NewRepository(db)
	service := category.NewService(repo)
	handler := category.NewHandler(service)

	r := router.Setup(handler)
	r.Run(":" + os.Getenv("PORT"))
}
