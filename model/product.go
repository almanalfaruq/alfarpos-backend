package model

type Product struct {
	Template
	Code       string   `gorm:"unique_index" json:"code" example:"unique-code"`
	Name       string   `json:"name" example:"product name"`
	BuyPrice   *int64   `json:"buy_price" example:"10000"`
	SellPrice  *int64   `json:"sell_price" example:"15000"`
	Quantity   *int64   `json:"quantity" example:"10"`
	CategoryID int64    `json:"category_id" example:"1"`
	Category   Category `gorm:"foreignkey:CategoryID" json:"category"`
	UnitID     int64    `json:"unit_id" example:"2"`
	Unit       Unit     `gorm:"foreignkey:UnitID" json:"unit"`
	ImageUrl   string   `json:"image_url" example:"http://localhost/image/image.jpg"`
	Discount   float32  `gorm:"default:0.00" json:"discount" example:"0.1"`
}
