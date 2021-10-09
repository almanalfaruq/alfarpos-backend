package stock

import (
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type StockRepository struct {
	db util.DBIface
}

func NewStockRepo(db util.DBIface) *StockRepository {
	return &StockRepository{
		db: db,
	}
}

func (repo *StockRepository) FindAll() ([]stockentity.Stock, error) {
	var stocks []stockentity.Stock
	db := repo.db.GetDb()
	return stocks, db.Set("gorm:auto_preload", true).Find(&stocks).Error
}

func (repo *StockRepository) FindByProduct(product productentity.Product) (stockentity.Stock, error) {
	var stock stockentity.Stock
	db := repo.db.GetDb()
	return stock, db.Set("gorm:auto_preload", true).Model(&product).Error
}

func (repo *StockRepository) New(stock stockentity.Stock) (stockentity.Stock, error) {
	db := repo.db.GetDb()
	return stock, db.Create(&stock).Error
}

func (repo *StockRepository) Update(stock stockentity.Stock) (stockentity.Stock, error) {
	var oldStock stockentity.Stock
	db := repo.db.GetDb()
	err := db.Where("id = ?", stock.ID).First(&oldStock).Error
	if err != nil {
		return stock, err
	}
	oldStock = stock
	return stock, db.Save(&oldStock).Error
}

func (repo *StockRepository) Delete(id int64) (stockentity.Stock, error) {
	var stock stockentity.Stock
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&stock)
	err := db.Delete(&stock).Error
	return stock, err
}

func (repo *StockRepository) DeleteAll() (int64, error) {
	var stock stockentity.Stock
	var stockCount int64
	db := repo.db.GetDb()
	err := db.Model(&stock).Count(&stockCount).Error
	if err != nil {
		return 0, err
	}
	return stockCount, db.Unscoped().Delete(&stock).Error
}
