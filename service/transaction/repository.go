package transaction

import (
	transactionentity "github.com/almanalfaruq/alfarpos-backend/model/transaction"
)

type moneyRepo interface {
	InsertMoneyTransaction(money transactionentity.Money) (transactionentity.Money, error)
	GetMoneyTransactionByFilter(status []int32, startDate, endDate, sort string) ([]transactionentity.Money, error)
}
