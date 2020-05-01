package controller

import (
	"mime/multipart"

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
	GetOneCustomer(id int) (model.Customer, error)
	GetAllCustomer() []model.Customer
	NewCustomer(customerData string) (model.Customer, error)
	UpdateCustomer(customerData string) (model.Customer, error)
	DeleteCustomer(id int) (model.Customer, error)
}

type orderDetailServiceIface interface {
	GetOrderDetailByOrder(orderDetailData string) ([]model.OrderDetail, error)
	DeleteOrderDetail(id int) (model.OrderDetail, error)
	DeleteOrderDetailByOrderId(orderDetailData string) (int, error)
}

type orderServiceIface interface {
	GetAllOrder() ([]model.Order, error)
	GetOneOrder(id int) (model.Order, error)
	GetOrderByInvoice(invoice string) (model.Order, error)
	GetOrderByUserId(userId int) ([]model.Order, error)
	NewOrder(OrderData string) (model.Order, error)
	UpdateOrder(OrderData string) (model.Order, error)
	DeleteOrder(id int) (model.Order, error)
}

type paymentServiceIface interface {
	GetAllPayment() ([]model.Payment, error)
	GetOnePayment(id int) (model.Payment, error)
	GetPaymentsByName(name string) ([]model.Payment, error)
	NewPayment(paymentData string) (model.Payment, error)
	UpdatePayment(paymentData string) (model.Payment, error)
	DeletePayment(id int) (model.Payment, error)
}

type printServiceIface interface {
	OrderByInvoiceToPdf(invoice string) *gofpdf.Fpdf
}

type productServiceIface interface {
	GetAllProduct() ([]model.Product, error)
	GetOneProduct(id int) (model.Product, error)
	GetOneProductByCode(code string) (model.Product, error)
	GetProductsByCode(productCode string) ([]model.Product, error)
	GetProductsByName(productName string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile multipart.File) error
	UpdateProduct(productData string) (model.Product, error)
	DeleteProduct(id int) (model.Product, error)
}

type stockServiceIface interface {
	GetByProduct(stockData string) (model.Stock, error)
	UpdateStock(stockData string) (model.Stock, error)
}

type unitServiceIface interface {
	GetAllUnit() ([]model.Unit, error)
	GetOneUnit(id int) (model.Unit, error)
	GetUnitsByName(name string) ([]model.Unit, error)
	NewUnit(unitData string) (model.Unit, error)
	UpdateUnit(unitData string) (model.Unit, error)
	DeleteUnit(id int) (model.Unit, error)
}

type userServiceIface interface {
	GetOneUser(id int) (model.User, error)
	GetAllUser() []model.User
	LoginUser(userData string) (string, error)
	NewUser(userData string) (model.User, error)
	UpdateUser(userData string) (model.User, error)
	DeleteUser(id int) (model.User, error)
}
