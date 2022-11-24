package models

type Cart struct {
	Id      int `json:"id"`
	ItemId  int `json:"item_id"`
	BuyerId int `json:"buyer_id"`
}
