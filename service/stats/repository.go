package stats

import (
	"context"
	"time"

	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	transactionentity "github.com/almanalfaruq/alfarpos-backend/model/transaction"
)

type moneyRepo interface {
	GetMoneyTransactionByFilter(status []int32, startDate, endDate, sort string) ([]transactionentity.Money, error)
}

type orderRepositoryIface interface {
	FindByFilter(status []int32, invoice, startDate, endDate, sort string, limit, page int32) ([]orderentity.Order, error)
}

type statsRepo interface {
	InsertShopStats(ctx context.Context, stats statsentity.ShopStats) (statsentity.ShopStats, error)
	GetByDate(ctx context.Context, date string) (statsentity.ShopStats, error)
	GetTrendingProducts(ctx context.Context, startDate, endDate time.Time, limit int) ([]statsentity.ProductStats, error)
	GetProductSales(ctx context.Context, productID int64, startDate, endDate time.Time) (statsentity.ProductStats, error)
	GetShopStatsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]statsentity.ShopStats, error)
}
