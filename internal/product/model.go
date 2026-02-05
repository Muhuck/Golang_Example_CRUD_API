package product

import (
	"crud-api/internal/category"
	"time"
)

type Product struct {
	ID         uint              `json:"id"`
	CategoryID uint              `json:"category_id"`
	Name       string            `json:"name"`
	Price      int               `json:"price"`
	Stock      int               `json:"stock"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Category   category.Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type ProductFilter struct {
	Name string
}
