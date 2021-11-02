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

func (r *MoneyRepo) GetMoneyTransactionByFilter(status []int32, startDate, endDate, sort string) ([]transaction.Money, error) {
	var monies []transaction.Money
	var whereClauses []string
	if len(status) > 0 {
		var statuses []string
		for _, st := range status {
			statuses = append(statuses, strconv.FormatInt(int64(st), 10))
		}
		whereClauses = append(whereClauses, fmt.Sprintf("money.type IN (%s)", strings.Join(statuses, ",")))
	}
	if startDate != "" || endDate != "" {
		if startDate == endDate {
			whereClauses = append(whereClauses, fmt.Sprintf("timezone('UTC', money.created_at)::date = '%s'", startDate))
		} else {
			if startDate != "" {
				whereClauses = append(whereClauses, fmt.Sprintf("timezone('UTC', money.created_at)::date>= '%s'", startDate))
			}
			if endDate != "" {
				whereClauses = append(whereClauses, fmt.Sprintf("timezone('UTC', money.created_at)::date <= '%s'", endDate))
			}
		}
	}
	sortSql := "money.id"
	if sort != "" {
		sortSql = fmt.Sprintf("%s %s", sortSql, sort)
	}
	db := r.db.GetDb()
	err := db.Where(strings.Join(whereClauses, " AND ")).Order(sortSql).Find(&monies).Error
	return monies, err
}
