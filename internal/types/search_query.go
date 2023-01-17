package types

type SearchQuery struct {
	Keyword string
	ShopId  string
	Offset  string `json:"offset" default:"0"`
	Limit   string `json:"limit" default:"10"`
}
