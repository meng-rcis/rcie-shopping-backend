package cartdto

type AddCartItemDTO struct {
	UserId    string `json:"user_id" validate:"required"`
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}
