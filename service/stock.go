package service

import (
	"encoding/json"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/repository"
)

type StockService struct {
	Product repository.IProductRepository
	Stock   repository.IStockRepository
}

type IStockService interface {
	GetByProduct(stockData string) (model.Stock, error)
	UpdateStock(stockData string) (model.Stock, error)
}

func (service *StockService) GetByProduct(stockData string) (model.Stock, error) {
	var stock model.Stock
	stockDataByte := []byte(stockData)
	err := json.Unmarshal(stockDataByte, &stock)
	if err != nil {
		return stock, err
	}
	product := service.Product.FindById(stock.ProductID)
	stock = service.Stock.FindByProduct(product)
	return stock, nil
}

func (service *StockService) UpdateStock(stockData string) (model.Stock, error) {
	var stock model.Stock
	stockDataByte := []byte(stockData)
	err := json.Unmarshal(stockDataByte, &stock)
	if err != nil {
		return stock, err
	}
	stock = service.Stock.Update(stock)
	return stock, nil
}
