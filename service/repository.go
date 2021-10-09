package service

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
)

//go:generate mockgen -source=repository.go -package=service -destination=repository_mock_test.go
type categoryRepositoryIface interface {
	FindAll() ([]model.Category, error)
	FindById(id int64) (model.Category, error)
	FindByName(name string) ([]model.Category, error)
	New(category model.Category) (model.Category, error)
	Update(category model.Category) (model.Category, error)
	Delete(id int64) (model.Category, error)
}

type customerRepositoryIface interface {
	FindAll() ([]model.Customer, error)
	FindById(id int64) (model.Customer, error)
	New(customer model.Customer) (model.Customer, error)
	Update(customer model.Customer) (model.Customer, error)
	Delete(id int64) (model.Customer, error)
}

type orderRepositoryIface interface {
	FindAll() []orderentity.Order
	FindById(id int64) (orderentity.Order, error)
	FindByInvoice(invoice string) (orderentity.Order, error)
	FindByUserId(userId int64) ([]orderentity.Order, error)
	FindByFilter(status []int32, invoice, startDate, endDate, sort string) ([]orderentity.Order, error)
	New(orderData orderentity.Order) (orderentity.Order, error)
	Update(orderData orderentity.Order) (orderentity.Order, error)
	UpdateStatus(orderID int64, status int32) (orderentity.Order, error)
	Delete(id int64) (orderentity.Order, error)
}

type paymentRepositoryIface interface {
	FindAll() ([]model.Payment, error)
	FindById(id int64) (model.Payment, error)
	FindByName(name string) ([]model.Payment, error)
	New(payment model.Payment) (model.Payment, error)
	Update(payment model.Payment) (model.Payment, error)
	Delete(id int64) (model.Payment, error)
}

type unitRepositoryIface interface {
	FindAll() ([]model.Unit, error)
	FindById(id int64) (model.Unit, error)
	FindByName(name string) ([]model.Unit, error)
	New(unit model.Unit) (model.Unit, error)
	Update(unit model.Unit) (model.Unit, error)
	Delete(id int64) (model.Unit, error)
}
