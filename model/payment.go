package model

type Payment struct {
	Template
	Name string `json:"name" example:"Cash"`
}
