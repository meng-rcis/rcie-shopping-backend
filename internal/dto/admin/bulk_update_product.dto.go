package admindto

type BulkUpdateProductDTO struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Status      string  `json:"status"`
}
