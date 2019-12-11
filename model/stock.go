package model

type Stock struct {
	Template
	ProductID int     `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
}
