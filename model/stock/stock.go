package stock

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/product"
)

type Stock struct {
	model.Template
	ProductID int64           `json:"product_id"`
	Product   product.Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int32           `json:"quantity"`
}
