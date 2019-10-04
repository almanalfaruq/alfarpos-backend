package model

type Stock struct {
	Template
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
}
