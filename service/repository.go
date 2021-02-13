package service

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
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

type orderDetailRepositoryIface interface {
	FindByOrder(order orderentity.Order) ([]model.OrderDetail, error)
	New(orderDetail model.OrderDetail) (model.OrderDetail, error)
	Update(orderDetail model.OrderDetail) (model.OrderDetail, error)
	Delete(id int64) (model.OrderDetail, error)
	DeleteByOrderId(id int64) (int64, error)
}

type orderRepositoryIface interface {
	FindAll() []orderentity.Order
	FindById(id int64) (orderentity.Order, error)
	FindByInvoice(invoice string) (orderentity.Order, error)
	FindByUserId(userId int64) ([]orderentity.Order, error)
	New(orderData orderentity.Order) (orderentity.Order, error)
	Update(orderData orderentity.Order) (orderentity.Order, error)
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

type productRepositoryIface interface {
	FindAll() ([]model.Product, error)
	FindById(id int64) (model.Product, error)
	FindByExactCode(code string) (model.Product, error)
	FindByCode(code string) ([]model.Product, error)
	FindByName(name string) ([]model.Product, error)
	FindByCategoryName(name string) ([]model.Product, error)
	FindByUnitName(name string) ([]model.Product, error)
	New(product model.Product) (model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(id int64) (model.Product, error)
	DeleteAll() (int64, error)
}

type stockRepositoryIface interface {
	FindAll() ([]model.Stock, error)
	FindByProduct(product model.Product) (model.Stock, error)
	New(stock model.Stock) (model.Stock, error)
	Update(stock model.Stock) (model.Stock, error)
	Delete(id int64) (model.Stock, error)
	DeleteAll() (int64, error)
}

type unitRepositoryIface interface {
	FindAll() ([]model.Unit, error)
	FindById(id int64) (model.Unit, error)
	FindByName(name string) ([]model.Unit, error)
	New(unit model.Unit) (model.Unit, error)
	Update(unit model.Unit) (model.Unit, error)
	Delete(id int64) (model.Unit, error)
}

type userRepositoryIface interface {
	FindAll() ([]userentity.User, error)
	FindById(id int64) (userentity.User, error)
	FindByUsername(username string) (userentity.User, error)
	FindByUsernameForLogin(username string) (userentity.User, error)
	New(userData userentity.User) (userentity.User, error)
	Update(userData userentity.User) (userentity.User, error)
	Delete(id int64) (userentity.User, error)
}
