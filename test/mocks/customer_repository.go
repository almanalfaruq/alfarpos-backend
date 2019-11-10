package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type CustomerRepository struct {
	mock.Mock
}

func (mock *CustomerRepository) FindById(id int) model.Customer {
	args := mock.Called(id)
	return args.Get(0).(model.Customer)
}

func (mock *CustomerRepository) FindAll() []model.Customer {
	args := mock.Called()
	return args.Get(0).([]model.Customer)
}

func (mock *CustomerRepository) New(category model.Customer) model.Customer {
	args := mock.Called(category)
	return args.Get(0).(model.Customer)
}

func (mock *CustomerRepository) Update(category model.Customer) model.Customer {
	args := mock.Called(category)
	return args.Get(0).(model.Customer)
}

func (mock *CustomerRepository) Delete(id int) model.Customer {
	args := mock.Called(id)
	return args.Get(0).(model.Customer)
}
