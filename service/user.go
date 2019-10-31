package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"../model"
	"../model/response"
	"../repository"
	"../util"
)

type UserService struct {
	User   repository.IUserRepository
	Config util.Config
}

type IUserService interface {
	GetOneUser(id int) (model.User, error)
	GetAllUser() []model.User
	LoginUser(userData string) (string, error)
	NewUser(userData string) (model.User, error)
	UpdateUser(userData string) (model.User, error)
	DeleteUser(id int) (model.User, error)
}

func (service *UserService) GetOneUser(id int) (model.User, error) {
	return service.User.FindById(id), nil
}

func (service *UserService) GetAllUser() []model.User {
	return service.User.FindAll()
}

func (service *UserService) LoginUser(userData string) (string, error) {
	var user model.User
	secretKey := []byte(service.Config.SecretKey)
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return "", err
	}
	userFromDb := service.User.FindByUsername(user.Username)
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("Username or Password mismatch")
	}
	user = userFromDb
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, response.TokenResponse{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix(),
			Issuer:    "AlfarPOS",
		},
	})
	tokenString, _ := token.SignedString(secretKey)
	return tokenString, nil
}

func (service *UserService) NewUser(userData string) (model.User, error) {
	var user model.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(encryptedPassword)
	user, err = service.User.New(user)
	if err != nil {
		return model.User{}, err
	}
	user.Password = ""
	return user, nil
}

func (service *UserService) UpdateUser(userData string) (model.User, error) {
	var user model.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	user = service.User.Update(user)
	return user, nil
}

func (service *UserService) DeleteUser(id int) (model.User, error) {
	return service.User.Delete(id), nil
}
