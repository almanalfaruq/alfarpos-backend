package repository

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	db dbIface
}

func NewOrderRepo(db dbIface) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repo *OrderRepository) FindAll() []model.Order {
	var orders []model.Order
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Find(&orders)
	return orders
}

func (repo *OrderRepository) FindById(id int64) (model.Order, error) {
	var order model.Order
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&order).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return model.Order{}, model.ErrNotFound
		}
		return model.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) FindByInvoice(invoice string) (model.Order, error) {
	var order model.Order
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("invoice = ?", invoice).First(&order).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return model.Order{}, model.ErrNotFound
		}
		return model.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) FindByUserId(userId int64) ([]model.Order, error) {
	var orders []model.Order
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []model.Order{}, model.ErrNotFound
		}
		return []model.Order{}, err
	}
	return orders, nil
}

func (repo *OrderRepository) New(order model.Order) (model.Order, error) {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(order)
	if !isNotExist {
		return order, fmt.Errorf("Order is exists")
	}
	db.Create(&order)
	db.Set("gorm:auto_preload", true).Where("id = ?", order.ID).First(&order)
	return order, nil
}

func (repo *OrderRepository) Update(order model.Order) (model.Order, error) {
	var oldOrder model.Order
	db := repo.db.GetDb()
	err := db.Where("id = ?", order.ID).First(&oldOrder).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return model.Order{}, model.ErrNotFound
		}
		return order, err
	}
	oldOrder = order
	db.Save(&oldOrder)
	return order, nil
}

func (repo *OrderRepository) Delete(id int64) (model.Order, error) {
	var order model.Order
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&order).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return model.Order{}, model.ErrNotFound
		}
		return model.Order{}, err
	}
	err = db.Delete(&order).Error
	return order, err
}
