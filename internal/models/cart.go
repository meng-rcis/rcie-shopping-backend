package models

import "time"

type CartItem struct {
	Id         string    `json:"id"`
	OwnerId    string    `json:"owner_id"`
	ProductId  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
