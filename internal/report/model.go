package report

type BestSellingProduct struct {
	Name       string `json:"name"`
	QtyTerjual int    `json:"qty_terjual"`
}

type DailyReport struct {
	TotalRevenue   int                `json:"total_revenue"`
	TotalTransaksi int                `json:"total_transaksi"`
	ProdukTerlaris BestSellingProduct `json:"produk_terlaris"`
}
