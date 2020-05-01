package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
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

func (repo *OrderRepository) FindById(id int) model.Order {
	var order model.Order
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&order)
	return order
}

func (repo *OrderRepository) FindByInvoice(invoice string) model.Order {
	var order model.Order
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("invoice = ?", invoice).First(&order)
	return order
}

func (repo *OrderRepository) FindByUserId(userId int) []model.Order {
	var orders []model.Order
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("user_id = ?", userId).Find(&orders)
	return orders
}

func (repo *OrderRepository) New(order model.Order) model.Order {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(order)
	if isNotExist {
		db.Create(&order)
	}
	db.Set("gorm:auto_preload", true).Where("id = ?", order.ID).First(&order)
	return order
}

func (repo *OrderRepository) Update(order model.Order) model.Order {
	var oldOrder model.Order
	db := repo.db.GetDb()
	db.Where("id = ?", order.ID).First(&oldOrder)
	oldOrder = order
	db.Save(&oldOrder)
	return order
}

func (repo *OrderRepository) Delete(id int) (model.Order, error) {
	var order model.Order
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&order)
	err := db.Delete(&order).Error
	return order, err
}
