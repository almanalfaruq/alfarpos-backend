package model

type OrderDetail struct {
	Template
	ProductOrder    `gorm:"embedded" json:"product_order"`
	Product         Product `gorm:"foreignkey:product_id" json:"product"`
	Quantity        int32   `json:"quantity" example:"2"`
	SubTotal        int64   `json:"sub_total" example:"5000"`
	OrderID         int64   `json:"order_id" example:"1"`
	UseSpecialPrice bool    `json:"use_special_price" example:"true"`
}
