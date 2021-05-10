package stats

import (
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type ShopStats struct {
	model.Template
	GrossProfit int64     `json:"gross_profit"`
	NetProfit   int64     `json:"net_profit"`
	OrderCount  int32     `json:"order_count"`
	SellAverage float64   `json:"sell_average"`
	MoneyIn     float64   `gorm:"type:decimal(10,2)" json:"money_in"`
	MoneyOut    float64   `gorm:"type:decimal(10,2)" json:"money_out"`
	Date        time.Time `gorm:"index" json:"date"`
}
