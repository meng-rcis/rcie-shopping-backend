package models

type Sample struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type NewSample struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreatedSample struct {
	Id int `json:"id"`
}

type UpdatedSample struct {
	Id      int  `json:"id"`
	Updated bool `json:"status"`
}

type DeleteSample struct {
	Id      int  `json:"id"`
	Deleted bool `json:"status"`
}
