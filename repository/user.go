package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/kataras/golog"
)

type UserRepository struct {
	db dbIface
}

func NewUserRepo(db dbIface) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindAll() []model.User {
	var users []model.User
	db := repo.db.GetDb()
	db.Find(&users)
	return users
}

func (repo *UserRepository) FindById(id int64) model.User {
	var user model.User
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&user)
	return user
}

func (repo *UserRepository) FindByUsername(username string) model.User {
	var user model.User
	db := repo.db.GetDb()
	db.Where("username = ?", username).First(&user)
	return user
}

func (repo *UserRepository) New(user model.User) (model.User, error) {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(user)
	if isNotExist {
		err := db.Create(&user).Error
		if err != nil {
			golog.Error(err)
			return user, err
		}
	}
	return user, nil
}

func (repo *UserRepository) Update(user model.User) model.User {
	var oldUser model.User
	db := repo.db.GetDb()
	db.Where("id = ?", user.ID).First(&oldUser)
	oldUser = user
	db.Save(&oldUser)
	return user
}

func (repo *UserRepository) Delete(id int64) model.User {
	var user model.User
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&user)
	db.Delete(&user)
	return user
}
