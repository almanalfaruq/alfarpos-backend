package service_test

import (
	"testing"

	. "github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/test/mocks"
	"github.com/almanalfaruq/alfarpos-backend/test/resources"
	"github.com/stretchr/testify/assert"
)

func TestStockGetByProduct(t *testing.T) {
	t.Run("GetByProduct - Success", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		stockRepository := new(mocks.StockRepository)

		productRepository.On("FindById", 1).Return(resources.Product1)
		stockRepository.On("FindByProduct", resources.Product1).Return(resources.Stock1)

		stockService := StockService{
			Product: productRepository,
			Stock:   stockRepository,
		}

		expectedResult := resources.Stock1

		jsonString := `{
			"product_id": 1
		}`

		actualResult, err := stockService.GetByProduct(jsonString)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("GetByProduct - Error", func(t *testing.T) {
		productRepository := new(mocks.ProductRepository)
		stockRepository := new(mocks.StockRepository)

		stockService := StockService{
			Product: productRepository,
			Stock:   stockRepository,
		}

		jsonString := `{
			product_id: 1
		}`

		actualResult, err := stockService.GetByProduct(jsonString)

		assert.NotNil(t, err)
		assert.Empty(t, actualResult)
	})
}
