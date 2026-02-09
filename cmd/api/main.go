package main

import (
	"crud-api/config"
	"crud-api/internal/category"
	"crud-api/internal/platform/database"
	"crud-api/internal/platform/router"
	"crud-api/internal/product"
	"crud-api/internal/report"
	"crud-api/internal/transaction"
	"os"
)

func main() {
	config.Load()

	db := database.Connect(os.Getenv("DATABASE_URL"))
	db.AutoMigrate(
		&category.Category{},
		&product.Product{},
		&transaction.Transaction{},
		&transaction.TransactionDetail{},
	)

	repo_category := category.NewRepository(db)
	service_category := category.NewService(repo_category)
	handler_category := category.NewHandler(service_category)

	repo_product := product.NewRepository(db)
	service_product := product.NewService(repo_product)
	handler_product := product.NewHandler(service_product)

	repo_transaction := transaction.NewRepository(db)
	service_transaction := transaction.NewService(repo_transaction, repo_product)
	handler_transaction := transaction.NewHandler(service_transaction)

	repo_report := report.NewRepository(db)
	service_report := report.NewService(repo_report)
	handler_report := report.NewHandler(service_report)

	r := router.Setup(handler_category, handler_product, handler_transaction, handler_report)
	r.Run(":" + os.Getenv("PORT"))
}
