package stats

import (
	"context"

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
}
