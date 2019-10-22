package model

type Product struct {
	Template
	Code       string   `gorm:"unique_index" json:"code"`
	Name       string   `json:"name"`
	BuyPrice   int      `json:"buy_price"`
	SellPrice  int      `json:"sell_price"`
	Quantity   int      `json:"quantity"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
	UnitID     int      `json:"unit_id"`
	Unit       Unit     `json:"unit"`
}
