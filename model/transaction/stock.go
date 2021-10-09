package transaction

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/product"
)

type Stock struct {
	model.Template
	Type      string          `db:"type" example:"in"`
	Product   product.Product `json:"product"`
	ProductID int64           `db:"product_id" json:"product_id" example:"29"`
	Qty       int32           `db:"qty" json:"qty" example:"10"`
}
