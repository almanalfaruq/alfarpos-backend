package transaction

import "github.com/almanalfaruq/alfarpos-backend/model/transaction"

type MoneyUsecase struct {
	repo moneyRepo
}

func New(repo moneyRepo) *MoneyUsecase {
	return &MoneyUsecase{
		repo: repo,
	}
}

func (u *MoneyUsecase) InsertMoney(money transaction.Money) (transaction.Money, error) {
	return u.repo.InsertMoneyTransaction(money)
}

func (u *MoneyUsecase) GetMoneyWithFilter(param transaction.GetMoneyWithFilterReq) ([]transaction.Money, error) {
	return u.repo.GetMoneyTransactionByFilter(param.Types, param.StartDate, param.EndDate, param.Sort)
}
