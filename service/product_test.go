package service_test

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/util"

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
		testName string
		args     func() (string, io.ReadCloser)
		mock     func() (*ProductService, error)
	}{
		{
			testName: "Error - Not Valid File",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("")
				return "", file
			},
			mock: func() (*ProductService, error) {
				return &ProductService{
					Product:  nil,
					Category: nil,
					Unit:     nil,
					Stock:    nil,
				}, errors.New("invalid argument")
			},
		},
		{
			testName: "Error - Sheet Not Exists",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "abcde", file
			},
			mock: func() (*ProductService, error) {
				return &ProductService{
					Product:  nil,
					Category: nil,
					Unit:     nil,
					Stock:    nil,
				}, errors.New("sheet abcde is not exist")
			},
		},
		{
			testName: "Success - No Import",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "", file
			},
			mock: func() (*ProductService, error) {
				return &ProductService{
					Product:  nil,
					Category: nil,
					Unit:     nil,
					Stock:    nil,
				}, nil
			},
		},
		{
			testName: "Success - No Blank Data - New Product",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "Product1", file
			},
			mock: func() (*ProductService, error) {
				productRepository := new(mocks.ProductRepository)
				categoryRepository := new(mocks.CategoryRepository)
				unitRepository := new(mocks.UnitRepository)
				stockRepository := new(mocks.StockRepository)

				categoryRepository.On("FindByName", "Category1").Return(resources.Categories[:1])
				unitRepository.On("FindByName", "Unit1").Return(resources.Units[:1])
				productRepository.On("FindByCode", "Product1").Return([]model.Product{})
				productStub := resources.Product1
				productStub.ID = 0
				productRepository.On("New", productStub).Return(resources.Product1)
				stockStub := resources.Stock1
				stockStub.ID = 0
				stockStub.Quantity = 0
				stockRepository.On("New", stockStub).Return(resources.Stock1)
				return &ProductService{
					Product:  productRepository,
					Category: categoryRepository,
					Unit:     unitRepository,
					Stock:    stockRepository,
				}, nil
			},
		},
		{
			testName: "Success - No Blank Data - New Category and Unit | Update Product",
			args: func() (string, io.ReadCloser) {
				file, _ := os.Open("../test/resources/test.xlsx")
				return "Product1", file
			},
			mock: func() (*ProductService, error) {
				productRepository := new(mocks.ProductRepository)
				categoryRepository := new(mocks.CategoryRepository)
				unitRepository := new(mocks.UnitRepository)
				stockRepository := new(mocks.StockRepository)

				categoryRepository.On("FindByName", "Category1").Return([]model.Category{})
				categoryStub := model.Category{Name: "Category1"}
				categoryRepository.On("New", categoryStub).Return(resources.Category1)
				unitRepository.On("FindByName", "Unit1").Return([]model.Unit{})
				unitStub := model.Unit{Name: "Unit1"}
				unitRepository.On("New", unitStub).Return(resources.Unit1)
				productRepository.On("FindByCode", "Product1").Return(resources.Products[:1])
				productRepository.On("Update", resources.Product1).Return(resources.Product1)
				stockStub := resources.Stock1
				stockStub.ID = 0
				stockStub.Quantity = 0
				stockRepository.On("New", stockStub).Return(resources.Stock1)
				return &ProductService{
					Product:  productRepository,
					Category: categoryRepository,
					Unit:     unitRepository,
					Stock:    stockRepository,
				}, nil
			},
		},
	}

	for _, tt := range testTable {
		sheetName, excelFile := tt.args()
		service, expectedResult := tt.mock()
		t.Run(tt.testName, func(t *testing.T) {
			actualResult := service.NewProductUsingExcel(sheetName, excelFile)
			assert.Equal(t, expectedResult, actualResult)
		})
		excelFile.Close()
	}
}

func TestProductUpdateProduct(t *testing.T) {
	productStub := resources.Product1
	productStub.SellPrice = util.ToInt64(55000)

	productRepository := new(mocks.ProductRepository)
	productRepository.On("Update", productStub).Return(productStub)

	testTable := []struct {
		testName string
		arg      func() string
		expect   model.Product
		wantErr  bool
	}{
		{
			testName: "Error - Failed Unmarshal JSON",
			arg: func() string {
				return `{product_id: 1}`
			},
			expect:  model.Product{},
			wantErr: true,
		},
		{
			testName: "Success",
			arg: func() string {
				jsonByte, _ := json.Marshal(productStub)
				return string(jsonByte)
			},
			expect: productStub,
		},
	}

	service := &ProductService{
		Product:  productRepository,
		Category: nil,
		Unit:     nil,
		Stock:    nil,
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			arg := tt.arg()
			actualResult, err := service.UpdateProduct(arg)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, actualResult)
		})
	}
}
