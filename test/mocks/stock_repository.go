package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type StockRepository struct {
	mock.Mock
}

func (mock *StockRepository) FindAll() []model.Stock {
	args := mock.Called()
	return args.Get(0).([]model.Stock)
}

func (mock *StockRepository) FindByProduct(product model.Product) model.Stock {
	args := mock.Called(product)
	return args.Get(0).(model.Stock)
}

func (mock *StockRepository) New(stock model.Stock) model.Stock {
	args := mock.Called(stock)
	return args.Get(0).(model.Stock)
}

func (mock *StockRepository) Update(stock model.Stock) model.Stock {
	args := mock.Called(stock)
	return args.Get(0).(model.Stock)
}

func (mock *StockRepository) Delete(id int) (model.Stock, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Stock), args.Get(1).(error)
}

func (mock *StockRepository) DeleteAll() int {
	args := mock.Called()
	return args.Get(0).(int)
}
