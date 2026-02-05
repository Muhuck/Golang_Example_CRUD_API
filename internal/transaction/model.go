package transaction

import (
	"crud-api/internal/product"
	"time"
)

type Transaction struct {
	ID        uint            `json:"id"`
	ProductID uint            `json:"product_id"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Product   product.Product `json:"product" gorm:"foreignKey:ProductID"`
}
