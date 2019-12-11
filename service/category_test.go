package service_test

import (
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/model"
	. "github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/test/mocks"
	"github.com/almanalfaruq/alfarpos-backend/test/resources"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryById(t *testing.T) {
	t.Run("Get Category By ID - Pass", func(t *testing.T) {
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("FindById", 1).Return(resources.Category1)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := resources.Category1

		actualResult, err := categoryService.GetOneCategory(1)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Get Category By ID - Error", func(t *testing.T) {
		var category model.Category
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("FindById", 5).Return(category)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := category

		actualResult, err := categoryService.GetOneCategory(5)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Category not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestGetAllCategory(t *testing.T) {
	categoryRepository := new(mocks.CategoryRepository)

	categoryRepository.On("FindAll").Return(resources.Categories)

	categoryService := CategoryService{
		ICategoryRepository: categoryRepository,
	}

	expectedResult := resources.Categories

	actualResult, err := categoryService.GetAllCategory()

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
	assert.Equal(t, expectedResult[2].ID, actualResult[2].ID)
	assert.Equal(t, expectedResult[2].Name, actualResult[2].Name)
	assert.Equal(t, expectedResult[3].ID, actualResult[3].ID)
	assert.Equal(t, expectedResult[3].Name, actualResult[3].Name)
}

func TestNewCategory(t *testing.T) {
	t.Run("New Category - Success", func(t *testing.T) {
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("New", resources.Category2).Return(resources.Category2)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := resources.Category2

		jsonCategory := `{"id": 2, "name": "Category2"}`

		actualResult, err := categoryService.NewCategory(jsonCategory)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("New Category - Error", func(t *testing.T) {
		var category model.Category
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("New", category).Return(category)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := category

		jsonCategory := `{name: "Category2"}`

		actualResult, err := categoryService.NewCategory(jsonCategory)

		assert.NotNil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("Update Category - Success", func(t *testing.T) {
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("Update", resources.Category2Updated).Return(resources.Category2Updated)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := resources.Category2Updated

		jsonCategory := `{"id": 2, "name": "Category2Updated"}`

		actualResult, err := categoryService.UpdateCategory(jsonCategory)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Update Category - Error", func(t *testing.T) {
		var category model.Category
		categoryRepository := new(mocks.CategoryRepository)

		categoryRepository.On("Update", category).Return(category)

		categoryService := CategoryService{
			ICategoryRepository: categoryRepository,
		}

		expectedResult := category

		jsonCategory := `{name: "Category2Updated"}`

		actualResult, err := categoryService.UpdateCategory(jsonCategory)

		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestDeleteCategory(t *testing.T) {
	categoryRepository := new(mocks.CategoryRepository)

	categoryRepository.On("Delete", 4).Return(resources.Category4, nil)

	categoryService := CategoryService{
		ICategoryRepository: categoryRepository,
	}

	expectedResult := resources.Category4

	id := 4

	actualResult, err := categoryService.DeleteCategory(id)

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult.ID, actualResult.ID)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
}
