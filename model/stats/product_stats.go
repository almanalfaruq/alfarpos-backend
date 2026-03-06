package stats

import (
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/product"
)

type ProductStats struct {
	model.Template
	ProductID    int64           `json:"product_id"`
	Product      product.Product `gorm:"foreignkey:ProductID" json:"product"`
	TotalSold    int64           `json:"total_sold"`
	TotalRevenue int64           `json:"total_revenue"`
	Date         time.Time       `gorm:"index" json:"date"`
}
