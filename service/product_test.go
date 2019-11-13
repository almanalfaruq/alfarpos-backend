package service_test

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/model"
	. "github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/test/mocks"
	"github.com/almanalfaruq/alfarpos-backend/test/resources"
	"github.com/stretchr/testify/assert"
)

func TestProductGetAllProduct(t *testing.T) {
	t.Run("GetAllProduct - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindAll").Return(resources.Products)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := resources.Products

		actualResult, err := productService.GetAllProduct()

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetOneProduct(t *testing.T) {
	t.Run("GetOneProduct - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindById", 2).Return(resources.Product2)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := resources.Product2

		actualResult, err := productService.GetOneProduct(2)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetOneProduct - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindById", 10).Return(model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := model.Product{}

		actualResult, err := productService.GetOneProduct(10)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Product not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetOneProductByCode(t *testing.T) {
	t.Run("GetOneProductByCode - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCode", "Product3").Return([]model.Product{
			resources.Product3,
		})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := resources.Product3

		actualResult, err := productService.GetOneProductByCode("Product3")

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetOneProductByCode - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCode", "Product10").Return([]model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := model.Product{}

		actualResult, err := productService.GetOneProductByCode("Product10")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Product not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetProductsByCode(t *testing.T) {
	t.Run("GetProductsByCode - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCode", "Product").Return(resources.Products)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := resources.Products

		actualResult, err := productService.GetProductsByCode("Product")

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetProductsByCode - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCode", "Productttt").Return([]model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{}

		actualResult, err := productService.GetProductsByCode("Productttt")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Products not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetProductsByName(t *testing.T) {
	t.Run("GetProductsByName - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByName", "product").Return(resources.Products)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := resources.Products

		actualResult, err := productService.GetProductsByName("Product")

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetProductsByName - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByName", "productttt").Return([]model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{}

		actualResult, err := productService.GetProductsByName("Productttt")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Products not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetProductsByCategoryName(t *testing.T) {
	t.Run("GetProductsByCategoryName - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCategoryName", "category3").Return([]model.Product{resources.Product3})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{
			resources.Product3,
		}

		actualResult, err := productService.GetProductsByCategoryName("Category3")

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetProductsByCategoryName - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByCategoryName", "category10").Return([]model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{}

		actualResult, err := productService.GetProductsByCategoryName("Category10")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Products not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductGetProductsByUnitName(t *testing.T) {
	t.Run("GetProductsByUnitName - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByUnitName", "unit2").Return([]model.Product{resources.Product2})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{
			resources.Product2,
		}

		actualResult, err := productService.GetProductsByUnitName("Unit2")

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetProductsByUnitName - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindByUnitName", "unit10").Return([]model.Product{})

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		expectedResult := []model.Product{}

		actualResult, err := productService.GetProductsByUnitName("Unit10")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Products not found")
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductNewProduct(t *testing.T) {
	t.Run("NewProduct - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("New", resources.Product3).Return(resources.Product3)
		stock3Edit := resources.Stock3
		stock3Edit.ID = 0
		stock3Edit.Quantity = 0
		stockRepository.On("New", stock3Edit).Return(stock3Edit)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		jsonBytes, _ := json.Marshal(resources.Product3)
		jsonString := string(jsonBytes)

		expectedResult := resources.Product3

		actualResult, err := productService.NewProduct(jsonString)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("NewProduct - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		categoryRepository := new(mocks.CategoryRepository)
		unitRepository := new(mocks.UnitRepository)
		stockRepository := new(mocks.StockRepository)

		productService := ProductService{
			Product:  productRepository,
			Category: categoryRepository,
			Unit:     unitRepository,
			Stock:    stockRepository,
		}

		jsonString := `{
			code: "product2"
		}`

		expectedResult := model.Product{}

		actualResult, err := productService.NewProduct(jsonString)

		assert.NotNil(t, err)
		assert.Empty(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestProductNewProductUsingExcel(t *testing.T) {
	testTable := []struct {
		testName           string
		sheetName          string
		excelFile          io.Reader
		categoryFindByName string
		categoryNew        model.Category
		unitFindByName     string
		unitNew            model.Unit
		productFindByCode  string
		productNew         model.Product
		productUpdate      model.Product
		stockNew           model.Stock
		expectedResult     error
	}{
		{"NewProductUsingExcel - Error", "", strings.NewReader(""), "", model.Category{}, "", model.Unit{}, "", model.Product{}, model.Product{}, model.Stock{}, errors.New("zip: not a valid zip file")},
	}
}
