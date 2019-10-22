package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type StockService struct {
	product repository.IProductRepository
	stock   repository.IStockRepository
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
	product := service.product.FindById(stock.ProductID)
	stock = service.stock.FindByProduct(product)
	return stock, nil
}

func (service *StockService) UpdateStock(stockData string) (model.Stock, error) {
	var stock model.Stock
	stockDataByte := []byte(stockData)
	err := json.Unmarshal(stockDataByte, &stock)
	if err != nil {
		return stock, err
	}
	stock = service.stock.Update(stock)
	return stock, nil
}
