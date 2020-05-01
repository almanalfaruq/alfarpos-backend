package repository

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type CategoryRepository struct {
	db dbIface
}

func NewCategoryRepo(db dbIface) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) FindAll() []model.Category {
	var categories []model.Category
	db := repo.db.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *CategoryRepository) FindById(id int64) (model.Category, error) {
	var category model.Category
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (repo *CategoryRepository) FindByName(name string) []model.Category {
	var categories []model.Category
	db := repo.db.GetDb()
	db.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&categories)
	return categories
}

func (repo *CategoryRepository) New(category model.Category) (model.Category, error) {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(category)
	if isNotExist {
		db.Create(&category)
		return category, nil
	}
	return category, fmt.Errorf("Category is exists")
}

func (repo *CategoryRepository) Update(category model.Category) (model.Category, error) {
	var oldCategory model.Category
	db := repo.db.GetDb()
	if err := db.Where("id = ?", category.ID).First(&oldCategory).Error; err != nil {
		return category, fmt.Errorf("Cannot update category as following: %v. Error: %v", category, err)
	}
	oldCategory = category
	db.Save(&oldCategory)
	return category, nil
}

func (repo *CategoryRepository) Delete(id int64) (model.Category, error) {
	var category model.Category
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&category)
	err := db.Delete(&category).Error
	return category, err
}
