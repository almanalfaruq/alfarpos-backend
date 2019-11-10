package model

type Product struct {
	Template
	Code       string   `gorm:"unique_index" json:"code"`
	Name       string   `json:"name"`
	BuyPrice   *int64   `json:"buy_price"`
	SellPrice  *int64   `json:"sell_price"`
	Quantity   *int64   `json:"quantity"`
	CategoryID int      `json:"category_id"`
	Category   Category `gorm:"foreignkey:CategoryID" json:"category"`
	UnitID     int      `json:"unit_id"`
	Unit       Unit     `gorm:"foreignkey:UnitID" json:"unit"`
	ImageUrl   string   `json:"image_url"`
	Discount   float32  `gorm:"default:0.00" json:"discount"`
}
