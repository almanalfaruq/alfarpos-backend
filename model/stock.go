package model

type Stock struct {
	Template
	ProductID int64   `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int64   `json:"quantity"`
}
