package mocks

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (mock *UserRepository) FindAll() []model.User {
	args := mock.Called()
	return args.Get(0).([]model.User)
}
func (mock *UserRepository) FindById(id int) model.User {
	args := mock.Called(id)
	return args.Get(0).(model.User)
}
func (mock *UserRepository) FindByUsername(username string) model.User {
	args := mock.Called(username)
	return args.Get(0).(model.User)
}
func (mock *UserRepository) New(user model.User) (model.User, error) {
	args := mock.Called(user)
	return args.Get(0).(model.User), args.Get(1).(error)
}
func (mock *UserRepository) Update(user model.User) model.User {
	args := mock.Called(user)
	return args.Get(0).(model.User)
}
func (mock *UserRepository) Delete(id int) model.User {
	args := mock.Called(id)
	return args.Get(0).(model.User)
}
