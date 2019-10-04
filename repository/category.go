package repository

import (
	"../model"
	"../util"
)

type CategoryRepository struct {
	util.IDatabaseConnection
}

type ICategoryRepository interface {
	FindById(id int) model.Category
	FindAll() []model.Category
	New(category model.Category) model.Category
	Update(category model.Category) model.Category
	Delete(id int) model.Category
}

func (repo *CategoryRepository) FindById(id int) model.Category {
	var category model.Category
	db := repo.GetDb()
	db.Where("id = ?", id).First(&category)
	return category
}

func (repo *CategoryRepository) FindAll() []model.Category {
	var categories []model.Category
	db := repo.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *CategoryRepository) New(category model.Category) model.Category {
	db := repo.GetDb()
	isNotExist := db.NewRecord(category)
	if isNotExist {
		db.Create(&category)
	}
	return category
}

func (repo *CategoryRepository) Update(category model.Category) model.Category {
	var oldCategory model.Category
	db := repo.GetDb()
	db.Where("id = ?", category.ID).First(&oldCategory)
	oldCategory = category
	db.Save(&oldCategory)
	return category
}

func (repo *CategoryRepository) Delete(id int) model.Category {
	var category model.Category
	db := repo.GetDb()
	db.Where("id = ?", id).First(&category)
	db.Delete(&category)
	return category
}
