package order

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
)

//go:generate mockgen -source=repository.go -package=order -destination=repository_mock_test.go
type orderDetailRepositoryIface interface {
	FindByOrder(ord orderentity.Order) ([]orderentity.OrderDetail, error)
	New(orderDetail orderentity.OrderDetail) (orderentity.OrderDetail, error)
	Update(orderDetail orderentity.OrderDetail) (orderentity.OrderDetail, error)
	Delete(id int64) (orderentity.OrderDetail, error)
	DeleteByOrderId(id int64) (int64, error)
}

type orderRepositoryIface interface {
	FindAll() []orderentity.Order
	FindById(id int64) (orderentity.Order, error)
	FindByInvoice(invoice string) (orderentity.Order, error)
	FindByUserId(userId int64) ([]orderentity.Order, error)
	FindByFilter(status []int32, invoice, startDate, endDate, sort string, limit, offset int32) ([]orderentity.Order, error)
	New(orderData orderentity.Order) (orderentity.Order, error)
	Update(orderData orderentity.Order) (orderentity.Order, error)
	UpdateStatus(orderID int64, status int32) (orderentity.Order, error)
	Delete(id int64) (orderentity.Order, error)
}

type paymentRepositoryIface interface {
	FindAll() ([]model.Payment, error)
	FindById(id int64) (model.Payment, error)
	FindByName(name string) ([]model.Payment, error)
	New(paymnt model.Payment) (model.Payment, error)
	Update(paymnt model.Payment) (model.Payment, error)
	Delete(id int64) (model.Payment, error)
}

type productRepositoryIface interface {
	FindAll() ([]productentity.Product, error)
	FindAllWithLimit(limit, offset int) ([]productentity.Product, error)
	FindById(id int64) (productentity.Product, error)
	FindByIDs(IDs []int64) ([]productentity.Product, error)
	FindByExactCode(code string) (productentity.Product, error)
	SearchBy(query string, limit, offset int) ([]productentity.Product, error)
	FindByCode(code string) ([]productentity.Product, error)
	FindByName(name string) ([]productentity.Product, error)
	FindByCategoryName(name string) ([]productentity.Product, error)
	FindByUnitName(name string) ([]productentity.Product, error)
	New(prd productentity.Product) (productentity.Product, error)
	Update(prd productentity.Product) (productentity.Product, error)
	Delete(id int64) (productentity.Product, error)
	DeleteAll() (int64, error)
}

type customerRepositoryIface interface {
	FindAll() ([]model.Customer, error)
	FindById(id int64) (model.Customer, error)
	New(cstm model.Customer) (model.Customer, error)
	Update(cstm model.Customer) (model.Customer, error)
	Delete(id int64) (model.Customer, error)
}
