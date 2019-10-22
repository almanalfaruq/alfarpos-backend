package repository

import (
	"../model"
	"../util"
)

type UserRepository struct {
	util.DatabaseConnection
}

type IUserRepository interface {
	FindAll() []model.User
	FindById(id int) model.User
	FindByUsername(username string) model.User
	Login(username string, password string) bool
	New(user model.User) model.User
	Update(user model.User) model.User
	Delete(id int) model.User
}

func (repo *UserRepository) FindAll() []model.User {
	var users []model.User
	db := repo.GetDb()
	db.Find(&users)
	return users
}

func (repo *UserRepository) FindById(id int) model.User {
	var user model.User
	db := repo.GetDb()
	db.Where("id = ?", id).First(&user)
	return user
}

func (repo *UserRepository) FindByUsername(username string) model.User {
	var user model.User
	db := repo.GetDb()
	db.Where("username = ?", username).First(&user)
	return user
}

func (repo *UserRepository) Login(username string, password string) bool {
	var user model.User
	db := repo.GetDb()
	db.Where("username = ? AND password = ?", username, password).First(&user)
	return user.Username == username && user.Password == password
}

func (repo *UserRepository) New(user model.User) model.User {
	db := repo.GetDb()
	isNotExist := db.NewRecord(user)
	if isNotExist {
		db.Create(&user)
	}
	return user
}

func (repo *UserRepository) Update(user model.User) model.User {
	var oldUser model.User
	db := repo.GetDb()
	db.Where("id = ?", user.ID).First(&oldUser)
	oldUser = user
	db.Save(&oldUser)
	return user
}

func (repo *UserRepository) Delete(id int) model.User {
	var user model.User
	db := repo.GetDb()
	db.Where("id = ?", id).First(&user)
	db.Delete(&user)
	return user
}
