package controller

import (
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
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
	GetAllCustomer() ([]model.Customer, error)
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
	GetAllOrder() ([]orderentity.Order, error)
	GetOneOrder(id int64) (orderentity.Order, error)
	GetOrderByInvoice(invoice string) (orderentity.Order, error)
	GetOrderByUserId(userId int64) ([]orderentity.Order, error)
	GetOrderUsingFilter(param orderentity.GetOrderUsingFilterParam) ([]orderentity.Order, error)
	NewOrder(orderData orderentity.Order) (orderentity.Order, error)
	UpdateOrder(orderData orderentity.Order) (orderentity.Order, error)
	UpdateOrderStatus(order orderentity.Order) (orderentity.Order, error)
	DeleteOrder(id int64) (orderentity.Order, error)
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
	GetAllProduct(limit, page int) (products []model.Product, hasNext bool, err error)
	GetOneProduct(id int64) (model.Product, error)
	GetProductsByIDs(IDs []int64) ([]model.Product, error)
	GetOneProductByCode(code string) (model.Product, error)
	GetProductsBySearchQuery(query string, limit, page int) (products []model.Product, hasNext bool, err error)
	GetProductsByCode(productCode string) ([]model.Product, error)
	GetProductsByName(productName string) ([]model.Product, error)
	GetProductsByCategoryName(categoryName string) ([]model.Product, error)
	GetProductsByUnitName(unitName string) ([]model.Product, error)
	ExportAllProductsToExcel() (*excelize.File, error)
	NewProduct(productData string) (model.Product, error)
	NewProductUsingExcel(sheetName string, excelFile io.Reader) (int, error)
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
	GetOneUser(id int64) (userentity.User, error)
	GetAllUser() ([]userentity.User, error)
	LoginUser(userData string) (userentity.UserResponse, error)
	NewUser(userData string) (userentity.User, error)
	UpdateUser(userData string) (userentity.User, error)
	DeleteUser(id int64) (userentity.User, error)
}
