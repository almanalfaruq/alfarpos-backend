package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/order"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	db dbIface
}

func NewOrderDetailRepo(db dbIface) *OrderDetailRepository {
	return &OrderDetailRepository{
		db: db,
	}
}

func (repo *OrderDetailRepository) FindByOrder(order order.Order) ([]model.OrderDetail, error) {
	var orderDetails []model.OrderDetail
	db := repo.db.GetDb()
	err := db.Joins("OrderDetail").Joins("Product").Model(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return []model.OrderDetail{}, model.ErrNotFound
		}
		return []model.OrderDetail{}, err
	}
	return orderDetails, nil
}

func (repo *OrderDetailRepository) New(orderDetail model.OrderDetail) (model.OrderDetail, error) {
	db := repo.db.GetDb()
	return orderDetail, db.Create(&orderDetail).Error
}

func (repo *OrderDetailRepository) Update(orderDetail model.OrderDetail) (model.OrderDetail, error) {
	var oldOrderDetail model.OrderDetail
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("id = ?", orderDetail.ID).First(&oldOrderDetail).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return model.OrderDetail{}, model.ErrNotFound
		}
		return model.OrderDetail{}, err
	}
	oldOrderDetail = orderDetail
	return orderDetail, db.Save(&oldOrderDetail).Error
}

func (repo *OrderDetailRepository) Delete(id int64) (model.OrderDetail, error) {
	var orderDetail model.OrderDetail
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&orderDetail).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return model.OrderDetail{}, model.ErrNotFound
		}
		return model.OrderDetail{}, err
	}
	return orderDetail, db.Delete(&orderDetail).Error
}

func (repo *OrderDetailRepository) DeleteByOrderId(id int64) (int64, error) {
	var orderDetailCount int64
	db := repo.db.GetDb()
	err := db.Model(&model.OrderDetail{}).Where("order_id = ?", id).Count(&orderDetailCount).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return 0, model.ErrNotFound
		}
		return 0, err
	}
	return orderDetailCount, db.Where("order_id = ?", id).Delete(&model.OrderDetail{}).Error
}
