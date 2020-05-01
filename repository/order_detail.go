package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

type OrderDetailRepository struct {
	db dbIface
}

func NewOrderDetailRepo(db dbIface) *OrderDetailRepository {
	return &OrderDetailRepository{
		db: db,
	}
}

func (repo *OrderDetailRepository) FindByOrder(order model.Order) []model.OrderDetail {
	var orderDetails []model.OrderDetail
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Model(&order).Related(&orderDetails)
	return orderDetails
}

func (repo *OrderDetailRepository) New(orderDetail model.OrderDetail) model.OrderDetail {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(orderDetail)
	if isNotExist {
		db.Create(&orderDetail)
	}
	return orderDetail
}

func (repo *OrderDetailRepository) Update(orderDetail model.OrderDetail) model.OrderDetail {
	var oldOrderDetail model.OrderDetail
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("id = ?", orderDetail.ID).First(&oldOrderDetail)
	oldOrderDetail = orderDetail
	db.Save(&oldOrderDetail)
	return orderDetail
}

func (repo *OrderDetailRepository) Delete(id int) model.OrderDetail {
	var orderDetail model.OrderDetail
	db := repo.db.GetDb()
	db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&orderDetail)
	db.Delete(&orderDetail)
	return orderDetail
}

func (repo *OrderDetailRepository) DeleteByOrderId(id int) int {
	var orderDetailCount int
	db := repo.db.GetDb()
	db.Model(&model.OrderDetail{}).Where("product_id = ?", id).Count(&orderDetailCount)
	db.Where("product_id = ?", id).Delete(&model.OrderDetail{})
	return orderDetailCount
}
