package transaction

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model/transaction"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type MoneyRepo struct {
	db util.DBIface
}

func NewMoney(db util.DBIface) *MoneyRepo {
	return &MoneyRepo{
		db: db,
	}
}

func (r *MoneyRepo) InsertMoneyTransaction(money transaction.Money) (transaction.Money, error) {
	db := r.db.GetDb()
	return money, db.Create(&money).Error
}

func (r *MoneyRepo) GetMoneyTransactionByFilter(status []int32, startDate, endDate string) ([]transaction.Money, error) {
	var monies []transaction.Money
	var whereClauses []string
	if len(status) > 0 {
		var statuses []string
		for _, st := range status {
			statuses = append(statuses, strconv.FormatInt(int64(st), 10))
		}
		whereClauses = append(whereClauses, fmt.Sprintf("money.status IN (%s)", strings.Join(statuses, ",")))
	}
	if startDate != "" || endDate != "" {
		if startDate == endDate {
			whereClauses = append(whereClauses, fmt.Sprintf("date(money.created_at) = '%s'", startDate))
		} else {
			if startDate != "" {
				whereClauses = append(whereClauses, fmt.Sprintf("money.created_at >= '%s'", startDate))
			}
			if endDate != "" {
				whereClauses = append(whereClauses, fmt.Sprintf("money.created_at <= '%s'", endDate))
			}
		}
	}
	db := r.db.GetDb()
	err := db.Where(strings.Join(whereClauses, " AND ")).Find(&monies).Error
	return monies, err
}
