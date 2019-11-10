package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type StockRepository struct {
	util.IDatabaseConnection
}

type IStockRepository interface {
	FindAll() []model.Stock
	FindByProduct(product model.Product) model.Stock
	New(stock model.Stock) model.Stock
	Update(stock model.Stock) model.Stock
	Delete(id int) (model.Stock, error)
	DeleteAll() int
}

func (repo *StockRepository) FindAll() []model.Stock {
	var stocks []model.Stock
	db := repo.GetDb()
	db.Set("gorm:auto_preload", true).Find(&stocks)
	return stocks
}

func (repo *StockRepository) FindByProduct(product model.Product) model.Stock {
	var stock model.Stock
	db := repo.GetDb()
	db.Set("gorm:auto_preload", true).Model(&product).Related(&stock)
	return stock
}

func (repo *StockRepository) New(stock model.Stock) model.Stock {
	db := repo.GetDb()
	isNotExist := db.NewRecord(stock)
	if isNotExist {
		db.Create(&stock)
	}
	return stock
}

func (repo *StockRepository) Update(stock model.Stock) model.Stock {
	var oldStock model.Stock
	db := repo.GetDb()
	db.Where("id = ?", stock.ID).First(&oldStock)
	oldStock = stock
	db.Save(&oldStock)
	return stock
}

func (repo *StockRepository) Delete(id int) (model.Stock, error) {
	var stock model.Stock
	db := repo.GetDb()
	db.Where("id = ?", id).First(&stock)
	err := db.Delete(&stock).Error
	return stock, err
}

func (repo *StockRepository) DeleteAll() int {
	var stock model.Stock
	var stockCount int
	db := repo.GetDb()
	db.Model(&stock).Count(&stockCount)
	db.Unscoped().Delete(&stock)
	return stockCount
}
