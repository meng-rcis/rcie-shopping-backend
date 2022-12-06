package cartdto

type UpdateCartItemDTO struct {
	Id       string `json:"id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}
