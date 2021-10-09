package stock

import (
	"encoding/json"

	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
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

func (service *StockService) GetByProduct(stockData string) (stockentity.Stock, error) {
	var stock stockentity.Stock
	stockDataByte := []byte(stockData)
	err := json.Unmarshal(stockDataByte, &stock)
	if err != nil {
		return stockentity.Stock{}, err
	}
	product, err := service.product.FindById(stock.ProductID)
	if err != nil {
		return stockentity.Stock{}, err
	}
	stock, err = service.stock.FindByProduct(product)
	if err != nil {
		return stockentity.Stock{}, err
	}
	return stock, nil
}

func (service *StockService) UpdateStock(stockData string) (stockentity.Stock, error) {
	var stock stockentity.Stock
	stockDataByte := []byte(stockData)
	err := json.Unmarshal(stockDataByte, &stock)
	if err != nil {
		return stock, err
	}
	return service.stock.Update(stock)
}
