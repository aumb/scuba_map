package models

type Location struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Status *bool  `json:"status,omitempty"`
}
