package models

import "time"

type Order struct {
	Id         string    `json:"id"`
	OwnerId    string    `json:"owner_id"`
	ProductId  string    `json:"product_id"`
	Status     string    `json:"status"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
