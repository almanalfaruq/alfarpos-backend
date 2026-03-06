package product

import (
	"database/sql"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/lib/pq"
)

type ProductV2 struct {
	model.Template
	Code            *string          `json:"code"`
	Name            string           `json:"name"`
	BuyPrice        *int64           `json:"buy_price"`
	SellPrice       *int64           `json:"sell_price"`
	Quantity        *int64           `json:"quantity"`
	CategoryID      int64            `json:"category_id"`
	Category        model.Category   `json:"category"`
	UnitID          int64            `json:"unit_id"`
	Unit            model.Unit       `json:"unit"`
	ImageUrl        string           `json:"image_url"`
	Discount        *float64         `json:"discount"`
	RelatedProducts pq.Int64Array    `json:"related_products"`
	ProductPrices   []ProductPriceV2 `json:"product_prices"`
	IsOpenPrice     bool             `json:"is_open_price"`
}

type ProductPriceV2 struct {
	model.Template
	ProductID          int64  `json:"product_id"`
	QuantityMultiplier int32  `json:"quantity_multiplier"`
	PricePerUnit       *int64 `json:"price_per_unit"`
	PricePerPacket     *int64 `json:"price_per_packet"`
}

func (p Product) ToV2() ProductV2 {
	return ProductV2{
		Template:        p.Template,
		Code:            nullStringToPointer(p.Code),
		Name:            p.Name,
		BuyPrice:        nullInt64ToPointer(p.BuyPrice),
		SellPrice:       nullInt64ToPointer(p.SellPrice),
		Quantity:        nullInt64ToPointer(p.Quantity),
		CategoryID:      p.CategoryID,
		Category:        p.Category,
		UnitID:          p.UnitID,
		Unit:            p.Unit,
		ImageUrl:        p.ImageUrl,
		Discount:        nullFloat64ToPointer(p.Discount),
		RelatedProducts: p.RelatedProducts,
		ProductPrices:   p.ProductPrices.ToV2(),
		IsOpenPrice:     p.IsOpenPrice,
	}
}

func (pp ProductPrices) ToV2() []ProductPriceV2 {
	var v2 []ProductPriceV2
	for _, p := range pp {
		v2 = append(v2, ProductPriceV2{
			Template:           p.Template,
			ProductID:          p.ProductID,
			QuantityMultiplier: p.QuantityMultiplier,
			PricePerUnit:       nullInt64ToPointer(p.PricePerUnit),
			PricePerPacket:     nullInt64ToPointer(p.PricePerPacket),
		})
	}
	return v2
}

func nullStringToPointer(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func nullInt64ToPointer(i sql.NullInt64) *int64 {
	if i.Valid {
		return &i.Int64
	}
	return nil
}

func nullFloat64ToPointer(f sql.NullFloat64) *float64 {
	if f.Valid {
		return &f.Float64
	}
	return nil
}
