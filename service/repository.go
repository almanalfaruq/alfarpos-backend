package service

import "github.com/almanalfaruq/alfarpos-backend/model"

//go:generate mockgen -source=repository.go -package=service -destination=repository_mock_test.go
type categoryRepositoryIface interface {
	FindAll() []model.Category
	FindById(id int64) (model.Category, error)
	FindByName(name string) []model.Category
	New(category model.Category) (model.Category, error)
	Update(category model.Category) (model.Category, error)
	Delete(id int64) (model.Category, error)
}

type customerRepositoryIface interface {
	FindAll() []model.Customer
	FindById(id int) model.Customer
	New(customer model.Customer) model.Customer
	Update(customer model.Customer) model.Customer
	Delete(id int) model.Customer
}

type orderDetailRepositoryIface interface {
	FindByOrder(order model.Order) []model.OrderDetail
	New(orderDetail model.OrderDetail) model.OrderDetail
	Update(orderDetail model.OrderDetail) model.OrderDetail
	Delete(id int) model.OrderDetail
	DeleteByOrderId(id int) int
}

type orderRepositoryIface interface {
	FindAll() []model.Order
	FindById(id int) model.Order
	FindByInvoice(invoice string) model.Order
	FindByUserId(userId int) []model.Order
	New(order model.Order) model.Order
	Update(order model.Order) model.Order
	Delete(id int) (model.Order, error)
}

type paymentRepositoryIface interface {
	FindAll() []model.Payment
	FindById(id int) model.Payment
	FindByName(name string) []model.Payment
	New(payment model.Payment) model.Payment
	Update(payment model.Payment) model.Payment
	Delete(id int) (model.Payment, error)
}

type productRepositoryIface interface {
	FindAll() []model.Product
	FindById(id int) model.Product
	FindByCode(code string) []model.Product
	FindByName(name string) []model.Product
	FindByCategoryName(name string) []model.Product
	FindByUnitName(name string) []model.Product
	New(product model.Product) model.Product
	Update(product model.Product) model.Product
	Delete(id int) (model.Product, error)
	DeleteAll() int
}

type stockRepositoryIface interface {
	FindAll() []model.Stock
	FindByProduct(product model.Product) model.Stock
	New(stock model.Stock) model.Stock
	Update(stock model.Stock) model.Stock
	Delete(id int) (model.Stock, error)
	DeleteAll() int
}

type unitRepositoryIface interface {
	FindAll() []model.Unit
	FindById(id int) model.Unit
	FindByName(name string) []model.Unit
	New(unit model.Unit) model.Unit
	Update(unit model.Unit) model.Unit
	Delete(id int) (model.Unit, error)
}

type userRepositoryIface interface {
	FindAll() []model.User
	FindById(id int) model.User
	FindByUsername(username string) model.User
	New(user model.User) (model.User, error)
	Update(user model.User) model.User
	Delete(id int) model.User
}
