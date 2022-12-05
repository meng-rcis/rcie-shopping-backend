package cartdto

type UpdateCartItemDTO struct {
	Id        string `json:"id" validate:"required"`
	UserId    string `json:"user_id" validate:"required"`
	ProductId string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}
