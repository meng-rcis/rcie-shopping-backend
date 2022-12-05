package admindto

type AddProductQuantityDTO struct {
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
	UserId    string `json:"user_id"`
}
