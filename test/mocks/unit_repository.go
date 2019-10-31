package mocks

import (
	"../../model"
	"github.com/stretchr/testify/mock"
)

type UnitRepository struct {
	mock.Mock
}

func (mock *UnitRepository) FindAll() []model.Unit {
	args := mock.Called()
	return args.Get(0).([]model.Unit)
}

func (mock *UnitRepository) FindById(id int) model.Unit {
	args := mock.Called(id)
	return args.Get(0).(model.Unit)
}

func (mock *UnitRepository) FindByName(name string) []model.Unit {
	args := mock.Called(name)
	return args.Get(0).([]model.Unit)
}

func (mock *UnitRepository) New(unit model.Unit) model.Unit {
	args := mock.Called(unit)
	return args.Get(0).(model.Unit)
}

func (mock *UnitRepository) Update(unit model.Unit) model.Unit {
	args := mock.Called(unit)
	return args.Get(0).(model.Unit)
}

func (mock *UnitRepository) Delete(id int) (model.Unit, error) {
	args := mock.Called(id)
	return args.Get(0).(model.Unit), args.Get(1).(error)
}
