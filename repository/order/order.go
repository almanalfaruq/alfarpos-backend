package order

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db util.DBIface
}

func NewOrderRepo(db util.DBIface) *OrderRepository {
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
	err := db.Preload("OrderDetails.Product").Preload("OrderDetails.Product.Unit").
		Preload("OrderDetails.Product.Category").Preload(clause.Associations).
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
	err := db.Preload("OrderDetails.Product").Preload("OrderDetails.Product.Unit").
		Preload("OrderDetails.Product.Category").Preload(clause.Associations).
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

func (repo *OrderRepository) FindByFilter(status []int32, invoice, startDate, endDate, sort string) ([]orderentity.Order, error) {
	var orders []orderentity.Order
	var whereClauses []string
	if invoice != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("orders.invoice = '%s'", invoice))
	}
	if len(status) > 0 {
		var statuses []string
		for _, st := range status {
			statuses = append(statuses, strconv.FormatInt(int64(st), 10))
		}
		whereClauses = append(whereClauses, fmt.Sprintf("orders.status IN (%s)", strings.Join(statuses, ",")))
	}
	if startDate != "" || endDate != "" {
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
	}
	sortSql := "orders.id"
	if sort != "" {
		sortSql = fmt.Sprintf("%s %s", sortSql, sort)
	}
	db := repo.db.GetDb()
	err := db.Preload("OrderDetails.Product").Preload("OrderDetails.Product.Unit").
		Preload("OrderDetails.Product.Category").Preload(clause.Associations).
		Where(strings.Join(whereClauses, " AND ")).Order(sortSql).Find(&orders).Error
	return orders, err
}

func (repo *OrderRepository) New(order orderentity.Order) (orderentity.Order, error) {
	db := repo.db.GetDb()
	err := db.Create(&order).Error
	if err != nil {
		return orderentity.Order{}, err
	}
	err = db.Preload("OrderDetails.Product").Preload("OrderDetails.Product.Unit").
		Preload("OrderDetails.Product.Category").Preload(clause.Associations).
		Where("id = ?", order.ID).
		First(&order).Error
	return order, err
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

func (repo *OrderRepository) UpdateStatus(orderID int64, status int32) (orderentity.Order, error) {
	var order orderentity.Order
	db := repo.db.GetDb()
	err := db.Where("id = ?", orderID).First(&order).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return orderentity.Order{}, model.ErrNotFound
		}
		return orderentity.Order{}, err
	}
	order.Status = status
	return order, db.Save(&order).Error
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
