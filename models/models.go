package models

type Suggest struct {
	Text string `json:"query"`
}

type FullSuggest struct {
	Id   int    `json:"id"`
	Text string `json:"query"`
}
