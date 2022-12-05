package admindto

type UpdateOrderStatusDTO struct {
	OrderId string `json:"order_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
	UserId  string `json:"user_id"`
}
