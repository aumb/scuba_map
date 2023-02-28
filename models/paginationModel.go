package models

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Offset int    `json:"offset"`
}
