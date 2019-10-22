package mocks

import (
	"../../model"
	"github.com/stretchr/testify/mock"
)

type PaymentRepository struct {
	mock.Mock
}

func (mock *PaymentRepository) FindById(id int) model.Payment {
	args := mock.Called(id)
	return args.Get(0).(model.Payment)
}

func (mock *PaymentRepository) FindAll() []model.Payment {
	args := mock.Called()
	return args.Get(0).([]model.Payment)
}

func (mock *PaymentRepository) New(payment model.Payment) model.Payment {
	args := mock.Called(payment)
	return args.Get(0).(model.Payment)
}

func (mock *PaymentRepository) Update(payment model.Payment) model.Payment {
	args := mock.Called(payment)
	return args.Get(0).(model.Payment)
}

func (mock *PaymentRepository) Delete(id int) model.Payment {
	args := mock.Called(id)
	return args.Get(0).(model.Payment)
}
