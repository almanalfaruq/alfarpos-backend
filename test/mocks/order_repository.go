package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type OrderRepository struct {
	mock.Mock
}

func (mock *OrderRepository) FindAll() []model.Order {
	args := mock.Called()
	return args.Get(0).([]model.Order)
}

func (mock *OrderRepository) FindById(id int) model.Order {
	args := mock.Called(id)
	return args.Get(0).(model.Order)
}

func (mock *OrderRepository) FindByInvoice(invoice string) model.Order {
	args := mock.Called(invoice)
	return args.Get(0).(model.Order)
}

func (mock *OrderRepository) FindByUserId(userId int) []model.Order {
	args := mock.Called(userId)
	return args.Get(0).([]model.Order)
}

func (mock *OrderRepository) New(order model.Order) model.Order {
	args := mock.Called(order)
	return args.Get(0).(model.Order)
}

func (mock *OrderRepository) Update(order model.Order) model.Order {
	args := mock.Called(order)
	return args.Get(0).(model.Order)
}

func (mock *OrderRepository) Delete(id int) (model.Order, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Order), args.Get(1).(error)
}
