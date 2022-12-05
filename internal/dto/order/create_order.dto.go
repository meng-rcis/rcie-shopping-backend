package orderdto

type CreateOrderDTO struct {
	UserId string `json:"user_id" validate:"required"`
	CartId string `json:"cart_id" validate:"required"`
}
