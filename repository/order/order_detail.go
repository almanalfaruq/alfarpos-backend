package order

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	db util.DBIface
}

func NewOrderDetailRepo(db util.DBIface) *OrderDetailRepository {
	return &OrderDetailRepository{
		db: db,
	}
}

func (repo *OrderDetailRepository) FindByOrder(order orderentity.Order) ([]orderentity.OrderDetail, error) {
	var orderDetails []orderentity.OrderDetail
	db := repo.db.GetDb()
	err := db.Joins("OrderDetail").Joins("Product").Model(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return []orderentity.OrderDetail{}, model.ErrNotFound
		}
		return []orderentity.OrderDetail{}, err
	}
	return orderDetails, nil
}

func (repo *OrderDetailRepository) New(orderDetail orderentity.OrderDetail) (orderentity.OrderDetail, error) {
	db := repo.db.GetDb()
	return orderDetail, db.Create(&orderDetail).Error
}

func (repo *OrderDetailRepository) Update(orderDetail orderentity.OrderDetail) (orderentity.OrderDetail, error) {
	var oldOrderDetail orderentity.OrderDetail
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("id = ?", orderDetail.ID).First(&oldOrderDetail).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.OrderDetail{}, model.ErrNotFound
		}
		return orderentity.OrderDetail{}, err
	}
	oldOrderDetail = orderDetail
	return orderDetail, db.Save(&oldOrderDetail).Error
}

func (repo *OrderDetailRepository) Delete(id int64) (orderentity.OrderDetail, error) {
	var orderDetail orderentity.OrderDetail
	db := repo.db.GetDb()
	err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&orderDetail).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.OrderDetail{}, model.ErrNotFound
		}
		return orderentity.OrderDetail{}, err
	}
	return orderDetail, db.Delete(&orderDetail).Error
}

func (repo *OrderDetailRepository) DeleteByOrderId(id int64) (int64, error) {
	var orderDetailCount int64
	db := repo.db.GetDb()
	err := db.Model(&orderentity.OrderDetail{}).Where("order_id = ?", id).Count(&orderDetailCount).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return 0, model.ErrNotFound
		}
		return 0, err
	}
	return orderDetailCount, db.Where("order_id = ?", id).Delete(&orderentity.OrderDetail{}).Error
}
