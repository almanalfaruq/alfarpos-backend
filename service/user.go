package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/auth"
)

type UserService struct {
	user userRepositoryIface
	conf util.Config
}

func NewUserService(conf util.Config, userRepo userRepositoryIface) *UserService {
	return &UserService{
		conf: conf,
		user: userRepo,
	}
}

func (service *UserService) GetOneUser(id int64) (model.User, error) {
	return service.user.FindById(id), nil
}

func (service *UserService) GetAllUser() []model.User {
	return service.user.FindAll()
}

func (service *UserService) LoginUser(userData string) (string, error) {
	var user model.User
	secretKey := []byte(service.conf.SecretKey)
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return "", err
	}
	userFromDb := service.user.FindByUsername(user.Username)
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("Username or Password mismatch")
	}
	user = userFromDb
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth.TokenData{
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
	user, err = service.user.New(user)
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
	user = service.user.Update(user)
	return user, nil
}

func (service *UserService) DeleteUser(id int64) (model.User, error) {
	return service.user.Delete(id), nil
}
