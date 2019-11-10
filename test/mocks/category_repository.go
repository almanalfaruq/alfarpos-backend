package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type CategoryRepository struct {
	mock.Mock
}

func (mock *CategoryRepository) FindAll() []model.Category {
	args := mock.Called()
	return args.Get(0).([]model.Category)
}

func (mock *CategoryRepository) FindById(id int) model.Category {
	args := mock.Called(id)
	return args.Get(0).(model.Category)
}

func (mock *CategoryRepository) FindByName(name string) []model.Category {
	args := mock.Called(name)
	return args.Get(0).([]model.Category)
}

func (mock *CategoryRepository) New(category model.Category) model.Category {
	args := mock.Called(category)
	return args.Get(0).(model.Category)
}

func (mock *CategoryRepository) Update(category model.Category) model.Category {
	args := mock.Called(category)
	return args.Get(0).(model.Category)
}

func (mock *CategoryRepository) Delete(id int) (model.Category, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Category), args.Get(1).(error)
}
