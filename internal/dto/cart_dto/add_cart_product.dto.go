package cart_dto

type AddCartProductDTO struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
