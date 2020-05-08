package model

type OrderDetail struct {
	Template
	ProductID int64   `json:"product_id" example:"1"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int64   `json:"quantity" example:"2"`
	SubTotal  int64   `json:"sub_total" example:"5000"`
	OrderID   int64   `json:"order_id" example:"1"`
	Order     Order   `gorm:"foreignkey:OrderID" json:"order"`
}
