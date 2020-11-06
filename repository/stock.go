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

func (repo *StockRepository) FindAll() ([]model.Stock, error) {
	var stocks []model.Stock
	db := repo.db.GetDb()
	return stocks, db.Set("gorm:auto_preload", true).Find(&stocks).Error
}

func (repo *StockRepository) FindByProduct(product model.Product) (model.Stock, error) {
	var stock model.Stock
	db := repo.db.GetDb()
	return stock, db.Set("gorm:auto_preload", true).Model(&product).Related(&stock).Error
}

func (repo *StockRepository) New(stock model.Stock) (model.Stock, error) {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(stock)
	if isNotExist {
		return stock, db.Create(&stock).Error
	}
	return stock, nil
}

func (repo *StockRepository) Update(stock model.Stock) (model.Stock, error) {
	var oldStock model.Stock
	db := repo.db.GetDb()
	err := db.Where("id = ?", stock.ID).First(&oldStock).Error
	if err != nil {
		return stock, err
	}
	oldStock = stock
	return stock, db.Save(&oldStock).Error
}

func (repo *StockRepository) Delete(id int64) (model.Stock, error) {
	var stock model.Stock
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&stock)
	err := db.Delete(&stock).Error
	return stock, err
}

func (repo *StockRepository) DeleteAll() (int64, error) {
	var stock model.Stock
	var stockCount int64
	db := repo.db.GetDb()
	err := db.Model(&stock).Count(&stockCount).Error
	if err != nil {
		return 0, err
	}
	return stockCount, db.Unscoped().Delete(&stock).Error
}
