package model

type OrderDetail struct {
	Template
	ProductID int     `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
	SubTotal  int     `json:"sub_total"`
	OrderID   int     `json:"order_id"`
	Order     Order   `gorm:"foreignkey:OrderID" json:"order"`
}
