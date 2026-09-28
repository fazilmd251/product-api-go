package models

import "time"

type Product struct {
	ProductID     int       `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Price         float64   `json:"price"`
	StockQuantity int       `json:"stock_quantity"`
	CategoryID    int       `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
}
