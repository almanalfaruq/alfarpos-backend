package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

type StockRepository struct {
	db dbIface
}

func NewStockRepo(db dbIface) *StockRepository {
	return &StockRepository{
		db: db,
	}
}

func (repo *StockRepository) FindAll() []model.Stock {
	var stocks []model.Stock
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Find(&stocks)
	return stocks
}

func (repo *StockRepository) FindByProduct(product model.Product) model.Stock {
	var stock model.Stock
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Model(&product).Related(&stock)
	return stock
}

func (repo *StockRepository) New(stock model.Stock) model.Stock {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(stock)
	if isNotExist {
		db.Create(&stock)
	}
	return stock
}

func (repo *StockRepository) Update(stock model.Stock) model.Stock {
	var oldStock model.Stock
	db := repo.db.GetDb()
	db.Where("id = ?", stock.ID).First(&oldStock)
	oldStock = stock
	db.Save(&oldStock)
	return stock
}

func (repo *StockRepository) Delete(id int64) (model.Stock, error) {
	var stock model.Stock
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&stock)
	err := db.Delete(&stock).Error
	return stock, err
}

func (repo *StockRepository) DeleteAll() int64 {
	var stock model.Stock
	var stockCount int64
	db := repo.db.GetDb()
	db.Model(&stock).Count(&stockCount)
	db.Unscoped().Delete(&stock)
	return stockCount
}
