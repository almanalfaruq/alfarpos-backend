package transaction

import "github.com/almanalfaruq/alfarpos-backend/model/transaction"

type moneyRepo interface {
	InsertMoneyTransaction(money transaction.Money) (transaction.Money, error)
	GetMoneyTransactionByFilter(status []int32, startDate, endDate string) ([]transaction.Money, error)
}
