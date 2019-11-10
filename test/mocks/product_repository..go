package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type ProductRepository struct {
	mock.Mock
}

func (mock *ProductRepository) FindAll() []model.Product {
	args := mock.Called()
	return args.Get(0).([]model.Product)
}

func (mock *ProductRepository) FindById(id int) model.Product {
	args := mock.Called(id)
	return args.Get(0).(model.Product)
}

func (mock *ProductRepository) FindByCode(code string) []model.Product {
	args := mock.Called(code)
	return args.Get(0).([]model.Product)
}

func (mock *ProductRepository) FindByName(name string) []model.Product {
	args := mock.Called(name)
	return args.Get(0).([]model.Product)
}

func (mock *ProductRepository) FindByCategoryName(name string) []model.Product {
	args := mock.Called(name)
	return args.Get(0).([]model.Product)
}

func (mock *ProductRepository) FindByUnitName(name string) []model.Product {
	args := mock.Called(name)
	return args.Get(0).([]model.Product)
}

func (mock *ProductRepository) New(product model.Product) model.Product {
	args := mock.Called(product)
	return args.Get(0).(model.Product)
}

func (mock *ProductRepository) Update(product model.Product) model.Product {
	args := mock.Called(product)
	return args.Get(0).(model.Product)
}

func (mock *ProductRepository) Delete(id int) (model.Product, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Product), args.Get(1).(error)
}

func (mock *ProductRepository) DeleteAll() int {
	args := mock.Called()
	return args.Get(0).(int)
}
