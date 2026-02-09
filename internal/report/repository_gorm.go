package report

import (
	"time"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db}
}

func (r *gormRepository) GetDailyReport(date time.Time) (*DailyReport, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	return r.GetReport(startOfDay, endOfDay)
}

func (r *gormRepository) GetReport(startDate, endDate time.Time) (*DailyReport, error) {
	var report DailyReport

	// Calculate Total Revenue and Total Transactions
	type RevenueResult struct {
		TotalRevenue   int
		TotalTransaksi int
	}
	var revenueResult RevenueResult

	// Assuming table name is transactions
	err := r.db.Table("transactions").
		Select("COALESCE(SUM(total_amount), 0) as total_revenue, COUNT(id) as total_transaksi").
		Where("created_at >= ? AND created_at < ?", startDate, endDate).
		Scan(&revenueResult).Error

	if err != nil {
		return nil, err
	}

	report.TotalRevenue = revenueResult.TotalRevenue
	report.TotalTransaksi = revenueResult.TotalTransaksi

	// Find Best Selling Product
	// Assuming table transaction_details stores product info
	type BestProductResult struct {
		ProductName string
		TotalQty    int
	}
	var bestProduct BestProductResult

	// Join transaction_details with transactions to filter by date
	err = r.db.Table("transaction_details").
		Select("transaction_details.product_name, SUM(transaction_details.quantity) as total_qty").
		Joins("JOIN transactions ON transactions.id = transaction_details.transaction_id").
		Where("transactions.created_at >= ? AND transactions.created_at < ?", startDate, endDate).
		Group("transaction_details.product_name").
		Order("total_qty DESC").
		Limit(1).
		Scan(&bestProduct).Error

	if err != nil {
		// Log error but maybe don't fail entire report?
		// For now, return error
		return nil, err
	}

	report.ProdukTerlaris = BestSellingProduct{
		Name:       bestProduct.ProductName,
		QtyTerjual: bestProduct.TotalQty,
	}

	return &report, nil
}
