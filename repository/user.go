package repository

import (
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
)

type UserRepository struct {
	db dbIface
}

func NewUserRepo(db dbIface) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindAll() ([]userentity.User, error) {
	var users []userentity.User
	db := repo.db.GetDb()
	return users, db.Find(&users).Error
}

func (repo *UserRepository) FindById(id int64) (userentity.User, error) {
	var user userentity.User
	db := repo.db.GetDb()
	return user, db.Where("id = ?", id).First(&user).Error
}

func (repo *UserRepository) FindByUsername(username string) (userentity.User, error) {
	var user userentity.User
	db := repo.db.GetDb()
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return userentity.User{}, err
	}
	user.Password = ""
	return user, err
}

func (repo *UserRepository) FindByUsernameForLogin(username string) (userentity.User, error) {
	var user userentity.User
	db := repo.db.GetDb()
	return user, db.Where("username = ?", username).First(&user).Error
}

func (repo *UserRepository) New(user userentity.User) (userentity.User, error) {
	db := repo.db.GetDb()
	return user, db.Create(&user).Error
}

func (repo *UserRepository) Update(user userentity.User) (userentity.User, error) {
	var oldUser userentity.User
	db := repo.db.GetDb()
	err := db.Where("id = ?", user.ID).First(&oldUser).Error
	if err != nil {
		return user, err
	}
	oldUser = user
	return user, db.Save(&oldUser).Error
}

func (repo *UserRepository) Delete(id int64) (userentity.User, error) {
	var user userentity.User
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, db.Delete(&user).Error
}
