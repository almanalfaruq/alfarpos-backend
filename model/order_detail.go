package model

type OrderDetail struct {
	Template
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	SubTotal  int     `json:"sub_total"`
	OrderID   int     `json:"order_id"`
	Order     Order   `json:"order"`
}
