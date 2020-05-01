package service

import (
	"encoding/json"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type StockService struct {
	product productRepositoryIface
	stock   stockRepositoryIface
}

func NewStockService(productRepo productRepositoryIface, stockRepo stockRepositoryIface) *StockService {
	return &StockService{
		product: productRepo,
		stock:   stockRepo,
	}
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
