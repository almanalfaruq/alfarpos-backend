package model

type OrderDetail struct {
	Template
	ProductID int64   `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int64   `json:"quantity"`
	SubTotal  int64   `json:"sub_total"`
	OrderID   int64   `json:"order_id"`
	Order     Order   `gorm:"foreignkey:OrderID" json:"order"`
}
