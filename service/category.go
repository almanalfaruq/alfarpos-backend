package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type CategoryService struct {
	repository.ICategoryRepository
}

type ICategoryService interface {
	GetAllCategory() []model.Category
	GetOneCategory(id int) (model.Category, error)
	NewCategory(categoryData string) (model.Category, error)
	UpdateCategory(categoryData string) (model.Category, error)
	DeleteCategory(id int) int
}

func (service *CategoryService) GetAllCategory() []model.Category {
	return service.FindAll()
}

func (service *CategoryService) GetOneCategory(id int) (model.Category, error) {
	return service.FindById(id), nil
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
	return service.Delete(id), nil
}
