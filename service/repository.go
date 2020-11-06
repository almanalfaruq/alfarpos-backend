package service

import "github.com/almanalfaruq/alfarpos-backend/model"

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
	FindAll() []model.Customer
	FindById(id int64) model.Customer
	New(customer model.Customer) model.Customer
	Update(customer model.Customer) model.Customer
	Delete(id int64) model.Customer
}

type orderDetailRepositoryIface interface {
	FindByOrder(order model.Order) ([]model.OrderDetail, error)
	New(orderDetail model.OrderDetail) (model.OrderDetail, error)
	Update(orderDetail model.OrderDetail) (model.OrderDetail, error)
	Delete(id int64) (model.OrderDetail, error)
	DeleteByOrderId(id int64) (int64, error)
}

type orderRepositoryIface interface {
	FindAll() []model.Order
	FindById(id int64) (model.Order, error)
	FindByInvoice(invoice string) (model.Order, error)
	FindByUserId(userId int64) ([]model.Order, error)
	New(order model.Order) (model.Order, error)
	Update(order model.Order) (model.Order, error)
	Delete(id int64) (model.Order, error)
}

type paymentRepositoryIface interface {
	FindAll() []model.Payment
	FindById(id int64) model.Payment
	FindByName(name string) []model.Payment
	New(payment model.Payment) model.Payment
	Update(payment model.Payment) model.Payment
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
	FindAll() []model.User
	FindById(id int64) model.User
	FindByUsername(username string) model.User
	New(user model.User) (model.User, error)
	Update(user model.User) model.User
	Delete(id int64) model.User
}
