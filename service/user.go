package service

import (
	"encoding/json"
	"errors"

	"../model"
	"../repository"
)

type UserService struct {
	repository.UserRepository
}

type IUserService interface {
	GetOneUser(id int) (model.User, error)
	GetAllUser() []model.User
	LoginUser(userData string) (model.User, error)
	NewUser(userData string) (model.User, error)
	UpdateUser(userData string) (model.User, error)
	DeleteUser(id int) (model.User, error)
}

func (service *UserService) GetOneUser(id int) (model.User, error) {
	return service.FindById(id), nil
}

func (service *UserService) GetAllUser() []model.User {
	return service.FindAll()
}

func (service *UserService) LoginUser(userData string) (model.User, error) {
	var user model.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		user = model.User{}
		return user, err
	}
	isLoggedIn := service.Login(user.Username, user.Password)
	if isLoggedIn {
		return service.FindByUsername(user.Username), nil
	}
	user = model.User{}
	return user, errors.New("Username or Password mismatch")
}

func (service *UserService) NewUser(userData string) (model.User, error) {
	var user model.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	return service.New(user), nil
}

func (service *UserService) UpdateUser(userData string) (model.User, error) {
	var user model.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	user = service.Update(user)
	return user, nil
}

func (service *UserService) DeleteUser(id int) (model.User, error) {
	return service.Delete(id), nil
}
