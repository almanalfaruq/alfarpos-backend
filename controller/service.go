package controller

import (
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
	GetAllCustomer() ([]model.Customer, error)
	NewCustomer(customerData string) (model.Customer, error)
	UpdateCustomer(customerData string) (model.Customer, error)
	DeleteCustomer(id int64) (model.Customer, error)
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

type unitServiceIface interface {
	GetAllUnit() ([]model.Unit, error)
	GetOneUnit(id int64) (model.Unit, error)
	GetUnitsByName(name string) ([]model.Unit, error)
	NewUnit(unitData string) (model.Unit, error)
	UpdateUnit(unitData string) (model.Unit, error)
	DeleteUnit(id int64) (model.Unit, error)
}
