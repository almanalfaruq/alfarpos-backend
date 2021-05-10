package stats

import (
	"context"

	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type StatsRepo struct {
	db util.DBIface
}

func NewStats(db util.DBIface) *StatsRepo {
	return &StatsRepo{
		db: db,
	}
}

func (r *StatsRepo) InsertShopStats(ctx context.Context, stats statsentity.ShopStats) (statsentity.ShopStats, error) {
	db := r.db.GetDb()
	return stats, db.Create(&stats).Error
}

func (r *StatsRepo) GetByDate(ctx context.Context, date string) (statsentity.ShopStats, error) {
	var result statsentity.ShopStats
	db := r.db.GetDb()
	return result, db.Where("date(shop_stats.date) = ?", date).First(&result).Error
}
