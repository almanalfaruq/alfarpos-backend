package product

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/lib/pq"
)

// Product field name
const (
	FieldCode          = "code"
	FieldName          = "name"
	FieldBuyPrice      = "buyprice"
	FieldSellPrice     = "sellprice"
	FieldQuantity      = "quantity"
	FieldCategory      = "category"
	FieldUnit          = "unit"
	FieldImageURL      = "imageurl"
	FieldDiscount      = "discount"
	FieldProductPrices = "productprices"
	FieldIsOpenPrice   = "isopenprice"
)

type Product struct {
	model.Template
	Code            sql.NullString  `gorm:"index" json:"code"`
	Name            string          `json:"name" example:"product name"`
	BuyPrice        sql.NullInt64   `json:"buy_price"`
	SellPrice       sql.NullInt64   `json:"sell_price"`
	Quantity        sql.NullInt64   `json:"quantity"`
	CategoryID      int64           `json:"category_id" example:"1"`
	Category        model.Category  `gorm:"foreignkey:CategoryID" json:"category"`
	UnitID          int64           `json:"unit_id" example:"2"`
	Unit            model.Unit      `gorm:"foreignkey:UnitID" json:"unit"`
	ImageUrl        string          `json:"image_url" example:"http://localhost/image/image.jpg"`
	Discount        sql.NullFloat64 `gorm:"default:0.00" json:"discount"`
	RelatedProducts pq.Int64Array   `gorm:"type:int8[]" json:"related_products"`
	ProductPrices   ProductPrices   `json:"product_prices"`
	IsOpenPrice     bool            `json:"is_open_price"`
}

type Products []Product

type ProductOrder struct {
	ProductID int64 `json:"product_id"`
	SellPrice int64 `json:"sell_price" example:"15000"`
}

type ProductPrice struct {
	model.Template
	ProductID          int64         `json:"product_id"`
	QuantityMultiplier int32         `json:"quantity_multiplier" example:"3"`
	PricePerUnit       sql.NullInt64 `json:"price_per_unit"`
	PricePerPacket     sql.NullInt64 `json:"price_per_packet"`
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
		PricePerUnit       sql.NullInt64 `json:"price_per_unit"`
		PricePerPacket     sql.NullInt64 `json:"price_per_packet"`
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

func DifferenceProductPrice(slice1 ProductPrices, slice2 ProductPrices) ProductPrices {
	var diff ProductPrices
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func (ps Products) Contains(product Product) (bool, int) {
	for idx, p := range ps {
		if p.Code.String == product.Code.String &&
			p.Name == product.Name &&
			p.Unit.Code == product.Unit.Code {
			return true, idx
		}
	}
	return false, -1
}

func (p Product) IsEqual(product Product) bool {
	return p.Code == product.Code &&
		p.Name == product.Name &&
		p.SellPrice.Int64 == product.SellPrice.Int64 &&
		p.BuyPrice.Int64 == product.BuyPrice.Int64 &&
		p.Unit.Code == product.Unit.Code
}

func (p *Product) ReplaceWith(product Product, fields ...string) {
	if len(fields) < 1 {
		p.Code = product.Code
		p.Name = product.Name
		p.BuyPrice = product.BuyPrice
		p.SellPrice = product.SellPrice
		p.Quantity = product.Quantity
		if product.CategoryID > 0 {
			p.CategoryID = product.CategoryID
		}
		p.Category = product.Category
		if product.UnitID > 0 {
			p.UnitID = product.UnitID
		}
		p.Unit = product.Unit
		p.ImageUrl = product.ImageUrl
		p.Discount = product.Discount
		p.ProductPrices = product.ProductPrices
		p.IsOpenPrice = product.IsOpenPrice
		return
	}

	for _, field := range fields {
		switch strings.ToLower(field) {
		case FieldCode:
			p.Code = product.Code
		case FieldName:
			p.Name = product.Name
		case FieldBuyPrice:
			p.BuyPrice = product.BuyPrice
		case FieldSellPrice:
			p.SellPrice = product.SellPrice
		case FieldQuantity:
			p.Quantity = product.Quantity
		case FieldCategory:
			if product.CategoryID > 0 {
				p.CategoryID = product.CategoryID
			}
			p.Category = product.Category
		case FieldUnit:
			if product.UnitID > 0 {
				p.UnitID = product.UnitID
			}
			p.Unit = product.Unit
		case FieldImageURL:
			p.ImageUrl = product.ImageUrl
		case FieldDiscount:
			p.Discount = product.Discount
		case FieldProductPrices:
			p.ProductPrices = product.ProductPrices
		case FieldIsOpenPrice:
			p.IsOpenPrice = product.IsOpenPrice
		}
	}
}
