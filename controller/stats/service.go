package stats

import (
	"context"

	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
)

type statsService interface {
	GetShopStats(ctx context.Context, date string) (statsentity.ShopStats, error)
	GetTrendingProducts(ctx context.Context, period string) ([]statsentity.ProductStats, error)
	GetProductSales(ctx context.Context, productID int64, period string) (statsentity.ProductStats, error)
	GetShopStatsByDateRange(ctx context.Context, startDate, endDate string) ([]statsentity.ShopStats, error)
}
