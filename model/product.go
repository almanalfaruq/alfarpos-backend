package model

import (
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

type Product struct {
	Template
	Code            sql.NullString  `gorm:"index" json:"code" example:"81921872917"`
	Name            string          `json:"name" example:"product name"`
	BuyPrice        sql.NullInt64   `json:"buy_price" example:"10000"`
	SellPrice       sql.NullInt64   `json:"sell_price" example:"15000"`
	Quantity        sql.NullInt64   `json:"quantity" example:"10"`
	CategoryID      int64           `json:"category_id" example:"1"`
	Category        Category        `gorm:"foreignkey:CategoryID" json:"category"`
	UnitID          int64           `json:"unit_id" example:"2"`
	Unit            Unit            `gorm:"foreignkey:UnitID" json:"unit"`
	ImageUrl        string          `json:"image_url" example:"http://localhost/image/image.jpg"`
	Discount        sql.NullFloat64 `gorm:"default:0.00" json:"discount" example:"0.1"`
	RelatedProducts pq.Int64Array   `gorm:"type:int8[]" json:"related_products"`
	ProductPrices   ProductPrices   `json:"product_prices"`
	IsOpenPrice     bool            `json:"is_open_price"`
}

type ProductOrder struct {
	ProductID int64 `json:"product_id"`
	SellPrice int64 `json:"sell_price" example:"15000"`
}

type ProductPrice struct {
	Template
	ProductID          int64         `json:"product_id"`
	QuantityMultiplier int32         `json:"quantity_multiplier" example:"3"`
	PricePerUnit       sql.NullInt64 `json:"price_per_unit" example:"15000"`
	PricePerPacket     sql.NullInt64 `json:"price_per_packet" example:"15000"`
}

type ProductPrices []ProductPrice

func (p *ProductPrices) ToString() (string, error) {
	if p == nil {
		return "", nil
	}

	if len(*p) == 0 {
		return "", nil
	}

	type productPricesWithoutTemplate struct {
		QuantityMultiplier int32         `json:"quantity_multiplier" example:"3"`
		PricePerUnit       sql.NullInt64 `json:"price_per_unit" example:"15000"`
		PricePerPacket     sql.NullInt64 `json:"price_per_packet" example:"15000"`
	}

	var pp []productPricesWithoutTemplate
	for _, productPrice := range *p {
		if productPrice.QuantityMultiplier == 0 || productPrice.QuantityMultiplier > 100 ||
			(productPrice.PricePerPacket.Int64 == 0 && productPrice.PricePerUnit.Int64 == 0) {
			return "", nil
		}
		pp = append(pp, productPricesWithoutTemplate{
			QuantityMultiplier: productPrice.QuantityMultiplier,
			PricePerUnit:       productPrice.PricePerUnit,
			PricePerPacket:     productPrice.PricePerPacket,
		})
	}

	byt, err := json.Marshal(pp)
	return string(byt), err
}

func (p ProductPrices) FromStringJson(s string) ProductPrices {
	var productPrices ProductPrices
	err := json.Unmarshal([]byte(s), &productPrices)
	if err != nil {
		return ProductPrices{}
	}

	return productPrices
}

type ProductResponse struct {
	Products []Product `json:"products"`
	HasNext  bool      `json:"has_next"`
}
