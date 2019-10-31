package service

import (
	"encoding/json"
	"errors"
	"strings"

	"../model"
	"../repository"
)

type CategoryService struct {
	repository.ICategoryRepository
}

type ICategoryService interface {
	GetAllCategory() ([]model.Category, error)
	GetOneCategory(id int) (model.Category, error)
	GetCategoriesByName(name string) ([]model.Category, error)
	NewCategory(categoryData string) (model.Category, error)
	UpdateCategory(categoryData string) (model.Category, error)
	DeleteCategory(id int) (model.Category, error)
}

func (service *CategoryService) GetAllCategory() ([]model.Category, error) {
	return service.FindAll(), nil
}

func (service *CategoryService) GetOneCategory(id int) (model.Category, error) {
	category := service.FindById(id)
	if category.ID == 0 {
		return category, errors.New("Category not found")
	}
	return category, nil
}

func (service *CategoryService) GetCategoriesByName(name string) ([]model.Category, error) {
	name = strings.ToLower(name)
	categories := service.FindByName(name)
	if len(categories) == 0 {
		return categories, errors.New("Categories not found")
	}
	return categories, nil
}

func (service *CategoryService) NewCategory(categoryData string) (model.Category, error) {
	var category model.Category
	categoryDataByte := []byte(categoryData)
	err := json.Unmarshal(categoryDataByte, &category)
	if err != nil {
		return category, err
	}
	return service.New(category), nil
}

func (service *CategoryService) UpdateCategory(categoryData string) (model.Category, error) {
	var category model.Category
	categoryDataByte := []byte(categoryData)
	err := json.Unmarshal(categoryDataByte, &category)
	if err != nil {
		return category, err
	}
	category = service.Update(category)
	return category, nil
}

func (service *CategoryService) DeleteCategory(id int) (model.Category, error) {
	return service.Delete(id)
}
