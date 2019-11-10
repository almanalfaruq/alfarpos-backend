package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type OrderDetailRepository struct {
	mock.Mock
}

func (mock *OrderDetailRepository) FindByOrder(order model.Order) []model.OrderDetail {
	args := mock.Called()
	return args.Get(0).([]model.OrderDetail)
}

func (mock *OrderDetailRepository) New(orderDetail model.OrderDetail) model.OrderDetail {
	args := mock.Called(orderDetail)
	return args.Get(0).(model.OrderDetail)
}

func (mock *OrderDetailRepository) Update(orderDetail model.OrderDetail) model.OrderDetail {
	args := mock.Called(orderDetail)
	return args.Get(0).(model.OrderDetail)
}

func (mock *OrderDetailRepository) Delete(id int) model.OrderDetail {
	args := mock.Called(id)
	return args.Get(0).(model.OrderDetail)
}

func (mock *OrderDetailRepository) DeleteByOrderId(id int) int {
	args := mock.Called(id)
	return args.Get(0).(int)
}
