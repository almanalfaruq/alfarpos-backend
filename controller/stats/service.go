package stats

import (
	"context"

	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
)

type statsService interface {
	GetShopStats(ctx context.Context, date string) (statsentity.ShopStats, error)
}
