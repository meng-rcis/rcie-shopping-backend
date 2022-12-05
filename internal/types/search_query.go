package types

type SearchQuery struct {
	Keyword string
	ShopId  string
	Page    string
	Limit   string `json:"limit" default:"10"`
}
