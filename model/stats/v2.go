package stats

import (
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/product"
)

type ShopStatsV2 struct {
	model.Template
	GrossProfit int64     `json:"gross_profit"`
	NetProfit   int64     `json:"net_profit"`
	OrderCount  int32     `json:"order_count"`
	SellAverage float64   `json:"sell_average"`
	MoneyIn     float64   `json:"money_in"`
	MoneyOut    float64   `json:"money_out"`
	Date        time.Time `json:"date"`
}

func (s ShopStats) ToV2() ShopStatsV2 {
	return ShopStatsV2{
		Template:    s.Template,
		GrossProfit: s.GrossProfit,
		NetProfit:   s.NetProfit,
		OrderCount:  s.OrderCount,
		SellAverage: s.SellAverage,
		MoneyIn:     s.MoneyIn,
		MoneyOut:    s.MoneyOut,
		Date:        s.Date,
	}
}

type ProductStatsV2 struct {
	model.Template
	ProductID    int64             `json:"product_id"`
	Product      product.ProductV2 `json:"product"`
	TotalSold    int64             `json:"total_sold"`
	TotalRevenue int64             `json:"total_revenue"`
	Date         time.Time         `json:"date"`
}

func (ps ProductStats) ToV2() ProductStatsV2 {
	return ProductStatsV2{
		Template:     ps.Template,
		ProductID:    ps.ProductID,
		Product:      ps.Product.ToV2(),
		TotalSold:    ps.TotalSold,
		TotalRevenue: ps.TotalRevenue,
		Date:         ps.Date,
	}
}
