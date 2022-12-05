package cartdto

type UpdateCartItemDTO struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
