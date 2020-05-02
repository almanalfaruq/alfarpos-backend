package controller

import (
	"io"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/jung-kurt/gofpdf"
)

//go:generate mockgen -source=service.go -package=controller -destination=service_mock_test.go
type categoryServiceIface interface {
	GetAllCategory() ([]model.Category, error)
	GetOneCategory(id int64) (model.Category, error)
	GetCategoriesByName(name string) ([]model.Category, error)
	NewCategory(name string) (model.Category, error)
	UpdateCategory(category model.Category) (model.Category, error)
	DeleteCategory(id int64) (model.Category, error)
}

type customerServiceIface interface {
	GetOneCustomer(id int64) (model.Customer, error)
	GetAllCustomer() []model.Customer
	NewCustomer(customerData string) (model.Customer, error)
	UpdateCustomer(customerData string) (model.Customer, error)
	DeleteCustomer(id int64) (model.Customer, error)
}

type orderDetailServiceIface interface {
	GetOrderDetailByOrder(orderDetailData string) ([]model.OrderDetail, error)
	DeleteOrderDetail(id int64) (model.OrderDetail, error)
	DeleteOrderDetailByOrderId(orderDetailData string) (int64, error)
}

type orderServiceIface interface {
	GetAllOrder() ([]model.Order, error)
	GetOneOrder(id int64) (model.Order, error)
	GetOrderByInvoice(invoice string) (model.Order, error)
	GetOrderByUserId(userId int64) ([]model.Order, error)
	NewOrder(order model.Order) (model.Order, error)
	UpdateOrder(order model.Order) (model.Order, error)
	DeleteOrder(id int64) (model.Order, error)
}

type paymentServiceIface interface {
	GetAllPayment() ([]model.Payment, error)
	GetOnePayment(id int64) (model.Payment, error)
	GetPaymentsByName(name string) ([]model.Payment, error)
	NewPayment(paymentData string) (model.Payment, error)
	UpdatePayment(paymentData string) (model.Payment, error)
	DeletePayment(id int64) (model.Payment, error)
}

type printServiceIface interface {
	OrderByInvoiceToPdf(invoice string) (*gofpdf.Fpdf, error)
}

type productServiceIface interface {
	GetAllProduct() ([]model.Product, error)
	GetOneProduct(id int64) (model.Product, error)
	GetOneProductByCode(code string) (model.Product, error)
	GetProductsByCode(productCode string) ([]model.Product, error)
	GetProductsByName(productName string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile io.Reader) error
	UpdateProduct(productData string) (model.Product, error)
	DeleteProduct(id int64) (model.Product, error)
}

type stockServiceIface interface {
	GetByProduct(stockData string) (model.Stock, error)
	UpdateStock(stockData string) (model.Stock, error)
}

type unitServiceIface interface {
	GetAllUnit() ([]model.Unit, error)
	GetOneUnit(id int64) (model.Unit, error)
	GetUnitsByName(name string) ([]model.Unit, error)
	NewUnit(unitData string) (model.Unit, error)
	UpdateUnit(unitData string) (model.Unit, error)
	DeleteUnit(id int64) (model.Unit, error)
}

type userServiceIface interface {
	GetOneUser(id int64) (model.User, error)
	GetAllUser() []model.User
	LoginUser(userData string) (string, error)
	NewUser(userData string) (model.User, error)
	UpdateUser(userData string) (model.User, error)
	DeleteUser(id int64) (model.User, error)
}
