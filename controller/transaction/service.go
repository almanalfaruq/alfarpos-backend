package transaction

import "github.com/almanalfaruq/alfarpos-backend/model/transaction"

type moneyUsecase interface {
	InsertMoney(money transaction.Money) (transaction.Money, error)
	GetMoneyWithFilter(param transaction.GetMoneyWithFilterReq) ([]transaction.Money, error)
}
