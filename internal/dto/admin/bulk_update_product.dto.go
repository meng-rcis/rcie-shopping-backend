package admindto

type BulkUpdateProductDTO struct {
	Description string  `json:"description"`
	Price       float64 `json:"price" min:"0"`
	Quantity    int     `json:"quantity" min:"0"`
	Status      string  `json:"status"`
}
