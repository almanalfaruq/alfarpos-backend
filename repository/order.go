package repository

import (
	"fmt"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db dbIface
}

func NewOrderRepo(db dbIface) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repo *OrderRepository) FindAll() []orderentity.Order {
	var orders []orderentity.Order
	db := repo.db.GetDb()
	db.Preload(clause.Associations).
		Find(&orders)
	return orders
}

func (repo *OrderRepository) FindById(id int64) (orderentity.Order, error) {
	var order orderentity.Order
	db := repo.db.GetDb()
	err := db.Preload(clause.Associations).
		Where("id = ?", id).
		First(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.Order{}, model.ErrNotFound
		}
		return orderentity.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) FindByInvoice(invoice string) (orderentity.Order, error) {
	var order orderentity.Order
	db := repo.db.GetDb()
	err := db.Preload(clause.Associations).
		Where("invoice = ?", invoice).
		First(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.Order{}, model.ErrNotFound
		}
		return orderentity.Order{}, err
	}
	return order, nil
}

func (repo *OrderRepository) FindByUserId(userId int64) ([]orderentity.Order, error) {
	var orders []orderentity.Order
	db := repo.db.GetDb()
	err := db.Preload(clause.Associations).
		Where("user_id = ?", userId).
		Find(&orders).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return []orderentity.Order{}, model.ErrNotFound
		}
		return []orderentity.Order{}, err
	}
	return orders, nil
}

func (repo *OrderRepository) FindByDate(startDate, endDate string) ([]orderentity.Order, error) {
	var orders []orderentity.Order
	var whereClauses []string
	if startDate == endDate {
		whereClauses = append(whereClauses, fmt.Sprintf("date(orders.created_at) = '%s'", startDate))
	} else {
		if startDate != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("orders.created_at >= '%s'", startDate))
		}
		if endDate != "" {
			whereClauses = append(whereClauses, fmt.Sprintf("orders.created_at <= '%s'", endDate))
		}
	}
	db := repo.db.GetDb()
	err := db.Preload(clause.Associations).
		Where(strings.Join(whereClauses, " AND ")).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) New(order orderentity.Order) (orderentity.Order, error) {
	db := repo.db.GetDb()
	return order, db.Create(&order).Error
}

func (repo *OrderRepository) Update(order orderentity.Order) (orderentity.Order, error) {
	var oldOrder orderentity.Order
	db := repo.db.GetDb()
	err := db.Where("id = ?", order.ID).First(&oldOrder).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.Order{}, model.ErrNotFound
		}
		return order, err
	}
	oldOrder = order
	db.Save(&oldOrder)
	return order, nil
}

func (repo *OrderRepository) Delete(id int64) (orderentity.Order, error) {
	var order orderentity.Order
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.Order{}, model.ErrNotFound
		}
		return orderentity.Order{}, err
	}
	err = db.Delete(&order).Error
	return order, err
}
