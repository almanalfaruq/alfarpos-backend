package repository

import (
	"../model"
	"../util"
)

type OrderDetailRepository struct {
	util.IDatabaseConnection
}

type IOrderDetailRepository interface {
	FindByOrder(order model.Order) []model.OrderDetail
	New(orderDetail model.OrderDetail) model.OrderDetail
	Update(orderDetail model.OrderDetail) model.OrderDetail
	Delete(id int) model.OrderDetail
	DeleteByOrderId(id int) int
}

func (repo *OrderDetailRepository) FindByOrder(order model.Order) []model.OrderDetail {
	var orderDetails []model.OrderDetail
	db := repo.GetDb()
	db.Model(&order).Related(&orderDetails)
	return orderDetails
}

func (repo *OrderDetailRepository) New(orderDetail model.OrderDetail) model.OrderDetail {
	db := repo.GetDb()
	isNotExist := db.NewRecord(orderDetail)
	if isNotExist {
		db.Create(&orderDetail)
	}
	return orderDetail
}

func (repo *OrderDetailRepository) Update(orderDetail model.OrderDetail) model.OrderDetail {
	var oldOrderDetail model.OrderDetail
	db := repo.GetDb()
	db.Where("id = ?", orderDetail.ID).First(&oldOrderDetail)
	oldOrderDetail = orderDetail
	db.Save(&oldOrderDetail)
	return orderDetail
}

func (repo *OrderDetailRepository) Delete(id int) model.OrderDetail {
	var orderDetail model.OrderDetail
	db := repo.GetDb()
	db.Where("id = ?", id).First(&orderDetail)
	db.Delete(&orderDetail)
	return orderDetail
}

func (repo *OrderDetailRepository) DeleteByOrderId(id int) int {
	var orderDetailCount int
	db := repo.GetDb()
	db.Model(&model.OrderDetail{}).Where("product_id = ?", id).Count(&orderDetailCount)
	db.Where("product_id = ?", id).Delete(&model.OrderDetail{})
	return orderDetailCount
}
