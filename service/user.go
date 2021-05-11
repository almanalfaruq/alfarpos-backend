package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
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

func (service *UserService) GetOneUser(id int64) (userentity.User, error) {
	return service.user.FindById(id)
}

func (service *UserService) GetAllUser() ([]userentity.User, error) {
	return service.user.FindAll()
}

func (service *UserService) LoginUser(userData string) (userentity.UserResponse, error) {
	var user userentity.User
	secretKey := []byte(service.conf.SecretKey)
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return userentity.UserResponse{}, err
	}
	userFromDb, err := service.user.FindByUsernameForLogin(user.Username)
	if err != nil {
		return userentity.UserResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err != nil {
		logger.Log.Error(err)
		return userentity.UserResponse{}, errors.New("Username or Password mismatch")
	}
	user = userFromDb
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, response.TokenData{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(24) * time.Hour).Unix(),
			Issuer:    "AlfarPOS",
		},
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return userentity.UserResponse{}, err
	}
	user.Password = ""
	return userentity.UserResponse{
		Token: tokenString,
		User:  user,
	}, nil
}

func (service *UserService) NewUser(userData string) (userentity.User, error) {
	var user userentity.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userentity.User{}, err
	}
	user.Password = string(encryptedPassword)
	user, err = service.user.New(user)
	if err != nil {
		return userentity.User{}, err
	}
	user.Password = ""
	return user, nil
}

func (service *UserService) UpdateUser(userData string) (userentity.User, error) {
	var user userentity.User
	userDataByte := []byte(userData)
	err := json.Unmarshal(userDataByte, &user)
	if err != nil {
		return user, err
	}
	return service.user.Update(user)
}

func (service *UserService) DeleteUser(id int64) (userentity.User, error) {
	return service.user.Delete(id)
}
