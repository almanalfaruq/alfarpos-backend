package stats

import (
	"context"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model/order"
	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
)

func (r *StatsRepo) GetTrendingProducts(ctx context.Context, startDate, endDate time.Time, limit int) ([]statsentity.ProductStats, error) {
	var results []statsentity.ProductStats
	db := r.db.GetDb()

	type Result struct {
		ProductID    int64 `json:"product_id"`
		TotalSold    int64 `json:"total_sold"`
		TotalRevenue int64 `json:"total_revenue"`
	}

	var rawResults []Result

	err := db.Table("order_details").
		Select("order_details.product_id, sum(order_details.quantity) as total_sold, sum(order_details.sub_total) as total_revenue").
		Joins("JOIN orders ON orders.id = order_details.order_id").
		Where("orders.status = ? AND orders.created_at BETWEEN ? AND ?", order.StatusFinish, startDate, endDate).
		Group("order_details.product_id").
		Order("total_sold DESC").
		Limit(limit).
		Scan(&rawResults).Error

	if err != nil {
		return nil, err
	}

	for _, res := range rawResults {
		results = append(results, statsentity.ProductStats{
			ProductID:    res.ProductID,
			TotalSold:    res.TotalSold,
			TotalRevenue: res.TotalRevenue,
		})
	}

	// Populate Product details
	for i := range results {
		db.Preload("Category").Preload("Unit").First(&results[i].Product, results[i].ProductID)
	}

	return results, nil
}

func (r *StatsRepo) GetProductSales(ctx context.Context, productID int64, startDate, endDate time.Time) (statsentity.ProductStats, error) {
	var result statsentity.ProductStats
	db := r.db.GetDb()

	type Result struct {
		TotalSold    int64 `json:"total_sold"`
		TotalRevenue int64 `json:"total_revenue"`
	}
	var rawResult Result

	err := db.Table("order_details").
		Select("sum(order_details.quantity) as total_sold, sum(order_details.sub_total) as total_revenue").
		Joins("JOIN orders ON orders.id = order_details.order_id").
		Where("orders.status = ? AND order_details.product_id = ? AND orders.created_at BETWEEN ? AND ?", order.StatusFinish, productID, startDate, endDate).
		Scan(&rawResult).Error

	if err != nil {
		return result, err
	}

	result.ProductID = productID
	result.TotalSold = rawResult.TotalSold
	result.TotalRevenue = rawResult.TotalRevenue

	// Populate Product details
	db.Preload("Category").Preload("Unit").First(&result.Product, productID)

	return result, nil
}
