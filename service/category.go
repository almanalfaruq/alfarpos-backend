package service

import (
	"errors"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type CategoryService struct {
	category categoryRepositoryIface
}

func NewCategoryService(categoryRepo categoryRepositoryIface) *CategoryService {
	return &CategoryService{
		category: categoryRepo,
	}
}

func (service *CategoryService) GetAllCategory() ([]model.Category, error) {
	return service.category.FindAll(), nil
}

func (service *CategoryService) GetOneCategory(id int64) (model.Category, error) {
	if id == 0 {
		return model.Category{}, ErrEmptyParam
	}
	category, err := service.category.FindById(id)
	if err != nil {
		return model.Category{}, err
	}
	if category.ID == 0 {
		return model.Category{}, errors.New("Category not found")
	}
	return category, nil
}

func (service *CategoryService) GetCategoriesByName(name string) ([]model.Category, error) {
	name = strings.ToLower(name)
	categories := service.category.FindByName(name)
	if len(categories) == 0 {
		return categories, errors.New("Categories not found")
	}
	return categories, nil
}

func (service *CategoryService) NewCategory(name string) (model.Category, error) {
	if name == "" {
		return model.Category{}, ErrEmptyParam
	}
	category := model.Category{
		Name: name,
	}
	return service.category.New(category)
}

func (service *CategoryService) UpdateCategory(category model.Category) (model.Category, error) {
	return service.category.Update(category)
}

func (service *CategoryService) DeleteCategory(id int64) (model.Category, error) {
	if id == 0 {
		return model.Category{}, ErrEmptyParam
	}
	return service.category.Delete(id)
}
