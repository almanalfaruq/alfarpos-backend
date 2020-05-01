// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package controller is a generated GoMock package.
package controller

import (
	model "github.com/almanalfaruq/alfarpos-backend/model"
	gomock "github.com/golang/mock/gomock"
	gofpdf "github.com/jung-kurt/gofpdf"
	multipart "mime/multipart"
	reflect "reflect"
)

// MockcategoryServiceIface is a mock of categoryServiceIface interface
type MockcategoryServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockcategoryServiceIfaceMockRecorder
}

// MockcategoryServiceIfaceMockRecorder is the mock recorder for MockcategoryServiceIface
type MockcategoryServiceIfaceMockRecorder struct {
	mock *MockcategoryServiceIface
}

// NewMockcategoryServiceIface creates a new mock instance
func NewMockcategoryServiceIface(ctrl *gomock.Controller) *MockcategoryServiceIface {
	mock := &MockcategoryServiceIface{ctrl: ctrl}
	mock.recorder = &MockcategoryServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcategoryServiceIface) EXPECT() *MockcategoryServiceIfaceMockRecorder {
	return m.recorder
}

// GetAllCategory mocks base method
func (m *MockcategoryServiceIface) GetAllCategory() ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategory")
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategory indicates an expected call of GetAllCategory
func (mr *MockcategoryServiceIfaceMockRecorder) GetAllCategory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategory", reflect.TypeOf((*MockcategoryServiceIface)(nil).GetAllCategory))
}

// GetOneCategory mocks base method
func (m *MockcategoryServiceIface) GetOneCategory(id int64) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneCategory", id)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneCategory indicates an expected call of GetOneCategory
func (mr *MockcategoryServiceIfaceMockRecorder) GetOneCategory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneCategory", reflect.TypeOf((*MockcategoryServiceIface)(nil).GetOneCategory), id)
}

// GetCategoriesByName mocks base method
func (m *MockcategoryServiceIface) GetCategoriesByName(name string) ([]model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesByName", name)
	ret0, _ := ret[0].([]model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesByName indicates an expected call of GetCategoriesByName
func (mr *MockcategoryServiceIfaceMockRecorder) GetCategoriesByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesByName", reflect.TypeOf((*MockcategoryServiceIface)(nil).GetCategoriesByName), name)
}

// NewCategory mocks base method
func (m *MockcategoryServiceIface) NewCategory(name string) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCategory", name)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCategory indicates an expected call of NewCategory
func (mr *MockcategoryServiceIfaceMockRecorder) NewCategory(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCategory", reflect.TypeOf((*MockcategoryServiceIface)(nil).NewCategory), name)
}

// UpdateCategory mocks base method
func (m *MockcategoryServiceIface) UpdateCategory(category model.Category) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", category)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory
func (mr *MockcategoryServiceIfaceMockRecorder) UpdateCategory(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockcategoryServiceIface)(nil).UpdateCategory), category)
}

// DeleteCategory mocks base method
func (m *MockcategoryServiceIface) DeleteCategory(id int64) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", id)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategory indicates an expected call of DeleteCategory
func (mr *MockcategoryServiceIfaceMockRecorder) DeleteCategory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockcategoryServiceIface)(nil).DeleteCategory), id)
}

// MockcustomerServiceIface is a mock of customerServiceIface interface
type MockcustomerServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockcustomerServiceIfaceMockRecorder
}

// MockcustomerServiceIfaceMockRecorder is the mock recorder for MockcustomerServiceIface
type MockcustomerServiceIfaceMockRecorder struct {
	mock *MockcustomerServiceIface
}

// NewMockcustomerServiceIface creates a new mock instance
func NewMockcustomerServiceIface(ctrl *gomock.Controller) *MockcustomerServiceIface {
	mock := &MockcustomerServiceIface{ctrl: ctrl}
	mock.recorder = &MockcustomerServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcustomerServiceIface) EXPECT() *MockcustomerServiceIfaceMockRecorder {
	return m.recorder
}

// GetOneCustomer mocks base method
func (m *MockcustomerServiceIface) GetOneCustomer(id int) (model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneCustomer", id)
	ret0, _ := ret[0].(model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneCustomer indicates an expected call of GetOneCustomer
func (mr *MockcustomerServiceIfaceMockRecorder) GetOneCustomer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneCustomer", reflect.TypeOf((*MockcustomerServiceIface)(nil).GetOneCustomer), id)
}

// GetAllCustomer mocks base method
func (m *MockcustomerServiceIface) GetAllCustomer() []model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomer")
	ret0, _ := ret[0].([]model.Customer)
	return ret0
}

// GetAllCustomer indicates an expected call of GetAllCustomer
func (mr *MockcustomerServiceIfaceMockRecorder) GetAllCustomer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomer", reflect.TypeOf((*MockcustomerServiceIface)(nil).GetAllCustomer))
}

// NewCustomer mocks base method
func (m *MockcustomerServiceIface) NewCustomer(customerData string) (model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCustomer", customerData)
	ret0, _ := ret[0].(model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCustomer indicates an expected call of NewCustomer
func (mr *MockcustomerServiceIfaceMockRecorder) NewCustomer(customerData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCustomer", reflect.TypeOf((*MockcustomerServiceIface)(nil).NewCustomer), customerData)
}

// UpdateCustomer mocks base method
func (m *MockcustomerServiceIface) UpdateCustomer(customerData string) (model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomer", customerData)
	ret0, _ := ret[0].(model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCustomer indicates an expected call of UpdateCustomer
func (mr *MockcustomerServiceIfaceMockRecorder) UpdateCustomer(customerData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomer", reflect.TypeOf((*MockcustomerServiceIface)(nil).UpdateCustomer), customerData)
}

// DeleteCustomer mocks base method
func (m *MockcustomerServiceIface) DeleteCustomer(id int) (model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCustomer", id)
	ret0, _ := ret[0].(model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCustomer indicates an expected call of DeleteCustomer
func (mr *MockcustomerServiceIfaceMockRecorder) DeleteCustomer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCustomer", reflect.TypeOf((*MockcustomerServiceIface)(nil).DeleteCustomer), id)
}

// MockorderDetailServiceIface is a mock of orderDetailServiceIface interface
type MockorderDetailServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockorderDetailServiceIfaceMockRecorder
}

// MockorderDetailServiceIfaceMockRecorder is the mock recorder for MockorderDetailServiceIface
type MockorderDetailServiceIfaceMockRecorder struct {
	mock *MockorderDetailServiceIface
}

// NewMockorderDetailServiceIface creates a new mock instance
func NewMockorderDetailServiceIface(ctrl *gomock.Controller) *MockorderDetailServiceIface {
	mock := &MockorderDetailServiceIface{ctrl: ctrl}
	mock.recorder = &MockorderDetailServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockorderDetailServiceIface) EXPECT() *MockorderDetailServiceIfaceMockRecorder {
	return m.recorder
}

// GetOrderDetailByOrder mocks base method
func (m *MockorderDetailServiceIface) GetOrderDetailByOrder(orderDetailData string) ([]model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetailByOrder", orderDetailData)
	ret0, _ := ret[0].([]model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetailByOrder indicates an expected call of GetOrderDetailByOrder
func (mr *MockorderDetailServiceIfaceMockRecorder) GetOrderDetailByOrder(orderDetailData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetailByOrder", reflect.TypeOf((*MockorderDetailServiceIface)(nil).GetOrderDetailByOrder), orderDetailData)
}

// DeleteOrderDetail mocks base method
func (m *MockorderDetailServiceIface) DeleteOrderDetail(id int) (model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderDetail", id)
	ret0, _ := ret[0].(model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrderDetail indicates an expected call of DeleteOrderDetail
func (mr *MockorderDetailServiceIfaceMockRecorder) DeleteOrderDetail(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderDetail", reflect.TypeOf((*MockorderDetailServiceIface)(nil).DeleteOrderDetail), id)
}

// DeleteOrderDetailByOrderId mocks base method
func (m *MockorderDetailServiceIface) DeleteOrderDetailByOrderId(orderDetailData string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderDetailByOrderId", orderDetailData)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrderDetailByOrderId indicates an expected call of DeleteOrderDetailByOrderId
func (mr *MockorderDetailServiceIfaceMockRecorder) DeleteOrderDetailByOrderId(orderDetailData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderDetailByOrderId", reflect.TypeOf((*MockorderDetailServiceIface)(nil).DeleteOrderDetailByOrderId), orderDetailData)
}

// MockorderServiceIface is a mock of orderServiceIface interface
type MockorderServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockorderServiceIfaceMockRecorder
}

// MockorderServiceIfaceMockRecorder is the mock recorder for MockorderServiceIface
type MockorderServiceIfaceMockRecorder struct {
	mock *MockorderServiceIface
}

// NewMockorderServiceIface creates a new mock instance
func NewMockorderServiceIface(ctrl *gomock.Controller) *MockorderServiceIface {
	mock := &MockorderServiceIface{ctrl: ctrl}
	mock.recorder = &MockorderServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockorderServiceIface) EXPECT() *MockorderServiceIfaceMockRecorder {
	return m.recorder
}

// GetAllOrder mocks base method
func (m *MockorderServiceIface) GetAllOrder() ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrder")
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrder indicates an expected call of GetAllOrder
func (mr *MockorderServiceIfaceMockRecorder) GetAllOrder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrder", reflect.TypeOf((*MockorderServiceIface)(nil).GetAllOrder))
}

// GetOneOrder mocks base method
func (m *MockorderServiceIface) GetOneOrder(id int) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneOrder", id)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneOrder indicates an expected call of GetOneOrder
func (mr *MockorderServiceIfaceMockRecorder) GetOneOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneOrder", reflect.TypeOf((*MockorderServiceIface)(nil).GetOneOrder), id)
}

// GetOrderByInvoice mocks base method
func (m *MockorderServiceIface) GetOrderByInvoice(invoice string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByInvoice", invoice)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByInvoice indicates an expected call of GetOrderByInvoice
func (mr *MockorderServiceIfaceMockRecorder) GetOrderByInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByInvoice", reflect.TypeOf((*MockorderServiceIface)(nil).GetOrderByInvoice), invoice)
}

// GetOrderByUserId mocks base method
func (m *MockorderServiceIface) GetOrderByUserId(userId int) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByUserId", userId)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByUserId indicates an expected call of GetOrderByUserId
func (mr *MockorderServiceIfaceMockRecorder) GetOrderByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByUserId", reflect.TypeOf((*MockorderServiceIface)(nil).GetOrderByUserId), userId)
}

// NewOrder mocks base method
func (m *MockorderServiceIface) NewOrder(OrderData string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewOrder", OrderData)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewOrder indicates an expected call of NewOrder
func (mr *MockorderServiceIfaceMockRecorder) NewOrder(OrderData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewOrder", reflect.TypeOf((*MockorderServiceIface)(nil).NewOrder), OrderData)
}

// UpdateOrder mocks base method
func (m *MockorderServiceIface) UpdateOrder(OrderData string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", OrderData)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder
func (mr *MockorderServiceIfaceMockRecorder) UpdateOrder(OrderData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockorderServiceIface)(nil).UpdateOrder), OrderData)
}

// DeleteOrder mocks base method
func (m *MockorderServiceIface) DeleteOrder(id int) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", id)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOrder indicates an expected call of DeleteOrder
func (mr *MockorderServiceIfaceMockRecorder) DeleteOrder(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockorderServiceIface)(nil).DeleteOrder), id)
}

// MockpaymentServiceIface is a mock of paymentServiceIface interface
type MockpaymentServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockpaymentServiceIfaceMockRecorder
}

// MockpaymentServiceIfaceMockRecorder is the mock recorder for MockpaymentServiceIface
type MockpaymentServiceIfaceMockRecorder struct {
	mock *MockpaymentServiceIface
}

// NewMockpaymentServiceIface creates a new mock instance
func NewMockpaymentServiceIface(ctrl *gomock.Controller) *MockpaymentServiceIface {
	mock := &MockpaymentServiceIface{ctrl: ctrl}
	mock.recorder = &MockpaymentServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockpaymentServiceIface) EXPECT() *MockpaymentServiceIfaceMockRecorder {
	return m.recorder
}

// GetAllPayment mocks base method
func (m *MockpaymentServiceIface) GetAllPayment() ([]model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPayment")
	ret0, _ := ret[0].([]model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPayment indicates an expected call of GetAllPayment
func (mr *MockpaymentServiceIfaceMockRecorder) GetAllPayment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPayment", reflect.TypeOf((*MockpaymentServiceIface)(nil).GetAllPayment))
}

// GetOnePayment mocks base method
func (m *MockpaymentServiceIface) GetOnePayment(id int) (model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnePayment", id)
	ret0, _ := ret[0].(model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOnePayment indicates an expected call of GetOnePayment
func (mr *MockpaymentServiceIfaceMockRecorder) GetOnePayment(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnePayment", reflect.TypeOf((*MockpaymentServiceIface)(nil).GetOnePayment), id)
}

// GetPaymentsByName mocks base method
func (m *MockpaymentServiceIface) GetPaymentsByName(name string) ([]model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentsByName", name)
	ret0, _ := ret[0].([]model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentsByName indicates an expected call of GetPaymentsByName
func (mr *MockpaymentServiceIfaceMockRecorder) GetPaymentsByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentsByName", reflect.TypeOf((*MockpaymentServiceIface)(nil).GetPaymentsByName), name)
}

// NewPayment mocks base method
func (m *MockpaymentServiceIface) NewPayment(paymentData string) (model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPayment", paymentData)
	ret0, _ := ret[0].(model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPayment indicates an expected call of NewPayment
func (mr *MockpaymentServiceIfaceMockRecorder) NewPayment(paymentData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPayment", reflect.TypeOf((*MockpaymentServiceIface)(nil).NewPayment), paymentData)
}

// UpdatePayment mocks base method
func (m *MockpaymentServiceIface) UpdatePayment(paymentData string) (model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", paymentData)
	ret0, _ := ret[0].(model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePayment indicates an expected call of UpdatePayment
func (mr *MockpaymentServiceIfaceMockRecorder) UpdatePayment(paymentData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockpaymentServiceIface)(nil).UpdatePayment), paymentData)
}

// DeletePayment mocks base method
func (m *MockpaymentServiceIface) DeletePayment(id int) (model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePayment", id)
	ret0, _ := ret[0].(model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePayment indicates an expected call of DeletePayment
func (mr *MockpaymentServiceIfaceMockRecorder) DeletePayment(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePayment", reflect.TypeOf((*MockpaymentServiceIface)(nil).DeletePayment), id)
}

// MockprintServiceIface is a mock of printServiceIface interface
type MockprintServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockprintServiceIfaceMockRecorder
}

// MockprintServiceIfaceMockRecorder is the mock recorder for MockprintServiceIface
type MockprintServiceIfaceMockRecorder struct {
	mock *MockprintServiceIface
}

// NewMockprintServiceIface creates a new mock instance
func NewMockprintServiceIface(ctrl *gomock.Controller) *MockprintServiceIface {
	mock := &MockprintServiceIface{ctrl: ctrl}
	mock.recorder = &MockprintServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprintServiceIface) EXPECT() *MockprintServiceIfaceMockRecorder {
	return m.recorder
}

// OrderByInvoiceToPdf mocks base method
func (m *MockprintServiceIface) OrderByInvoiceToPdf(invoice string) *gofpdf.Fpdf {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderByInvoiceToPdf", invoice)
	ret0, _ := ret[0].(*gofpdf.Fpdf)
	return ret0
}

// OrderByInvoiceToPdf indicates an expected call of OrderByInvoiceToPdf
func (mr *MockprintServiceIfaceMockRecorder) OrderByInvoiceToPdf(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderByInvoiceToPdf", reflect.TypeOf((*MockprintServiceIface)(nil).OrderByInvoiceToPdf), invoice)
}

// MockproductServiceIface is a mock of productServiceIface interface
type MockproductServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockproductServiceIfaceMockRecorder
}

// MockproductServiceIfaceMockRecorder is the mock recorder for MockproductServiceIface
type MockproductServiceIfaceMockRecorder struct {
	mock *MockproductServiceIface
}

// NewMockproductServiceIface creates a new mock instance
func NewMockproductServiceIface(ctrl *gomock.Controller) *MockproductServiceIface {
	mock := &MockproductServiceIface{ctrl: ctrl}
	mock.recorder = &MockproductServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockproductServiceIface) EXPECT() *MockproductServiceIfaceMockRecorder {
	return m.recorder
}

// GetAllProduct mocks base method
func (m *MockproductServiceIface) GetAllProduct() ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProduct")
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProduct indicates an expected call of GetAllProduct
func (mr *MockproductServiceIfaceMockRecorder) GetAllProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProduct", reflect.TypeOf((*MockproductServiceIface)(nil).GetAllProduct))
}

// GetOneProduct mocks base method
func (m *MockproductServiceIface) GetOneProduct(id int) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneProduct", id)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneProduct indicates an expected call of GetOneProduct
func (mr *MockproductServiceIfaceMockRecorder) GetOneProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneProduct", reflect.TypeOf((*MockproductServiceIface)(nil).GetOneProduct), id)
}

// GetOneProductByCode mocks base method
func (m *MockproductServiceIface) GetOneProductByCode(code string) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneProductByCode", code)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneProductByCode indicates an expected call of GetOneProductByCode
func (mr *MockproductServiceIfaceMockRecorder) GetOneProductByCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneProductByCode", reflect.TypeOf((*MockproductServiceIface)(nil).GetOneProductByCode), code)
}

// GetProductsByCode mocks base method
func (m *MockproductServiceIface) GetProductsByCode(productCode string) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByCode", productCode)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByCode indicates an expected call of GetProductsByCode
func (mr *MockproductServiceIfaceMockRecorder) GetProductsByCode(productCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCode", reflect.TypeOf((*MockproductServiceIface)(nil).GetProductsByCode), productCode)
}

// GetProductsByName mocks base method
func (m *MockproductServiceIface) GetProductsByName(productName string) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByName", productName)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByName indicates an expected call of GetProductsByName
func (mr *MockproductServiceIfaceMockRecorder) GetProductsByName(productName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByName", reflect.TypeOf((*MockproductServiceIface)(nil).GetProductsByName), productName)
}

// GetProductsByCategoryName mocks base method
func (m *MockproductServiceIface) GetProductsByCategoryName(categoryName string) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByCategoryName", categoryName)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByCategoryName indicates an expected call of GetProductsByCategoryName
func (mr *MockproductServiceIfaceMockRecorder) GetProductsByCategoryName(categoryName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCategoryName", reflect.TypeOf((*MockproductServiceIface)(nil).GetProductsByCategoryName), categoryName)
}

// GetProductsByUnitName mocks base method
func (m *MockproductServiceIface) GetProductsByUnitName(unitName string) ([]model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByUnitName", unitName)
	ret0, _ := ret[0].([]model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByUnitName indicates an expected call of GetProductsByUnitName
func (mr *MockproductServiceIfaceMockRecorder) GetProductsByUnitName(unitName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByUnitName", reflect.TypeOf((*MockproductServiceIface)(nil).GetProductsByUnitName), unitName)
}

// NewProduct mocks base method
func (m *MockproductServiceIface) NewProduct(productData string) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProduct", productData)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProduct indicates an expected call of NewProduct
func (mr *MockproductServiceIfaceMockRecorder) NewProduct(productData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProduct", reflect.TypeOf((*MockproductServiceIface)(nil).NewProduct), productData)
}

// NewProductUsingExcel mocks base method
func (m *MockproductServiceIface) NewProductUsingExcel(sheetName string, excelFile multipart.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProductUsingExcel", sheetName, excelFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewProductUsingExcel indicates an expected call of NewProductUsingExcel
func (mr *MockproductServiceIfaceMockRecorder) NewProductUsingExcel(sheetName, excelFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProductUsingExcel", reflect.TypeOf((*MockproductServiceIface)(nil).NewProductUsingExcel), sheetName, excelFile)
}

// UpdateProduct mocks base method
func (m *MockproductServiceIface) UpdateProduct(productData string) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", productData)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct
func (mr *MockproductServiceIfaceMockRecorder) UpdateProduct(productData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockproductServiceIface)(nil).UpdateProduct), productData)
}

// DeleteProduct mocks base method
func (m *MockproductServiceIface) DeleteProduct(id int) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", id)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct
func (mr *MockproductServiceIfaceMockRecorder) DeleteProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockproductServiceIface)(nil).DeleteProduct), id)
}

// MockstockServiceIface is a mock of stockServiceIface interface
type MockstockServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockstockServiceIfaceMockRecorder
}

// MockstockServiceIfaceMockRecorder is the mock recorder for MockstockServiceIface
type MockstockServiceIfaceMockRecorder struct {
	mock *MockstockServiceIface
}

// NewMockstockServiceIface creates a new mock instance
func NewMockstockServiceIface(ctrl *gomock.Controller) *MockstockServiceIface {
	mock := &MockstockServiceIface{ctrl: ctrl}
	mock.recorder = &MockstockServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstockServiceIface) EXPECT() *MockstockServiceIfaceMockRecorder {
	return m.recorder
}

// GetByProduct mocks base method
func (m *MockstockServiceIface) GetByProduct(stockData string) (model.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByProduct", stockData)
	ret0, _ := ret[0].(model.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByProduct indicates an expected call of GetByProduct
func (mr *MockstockServiceIfaceMockRecorder) GetByProduct(stockData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByProduct", reflect.TypeOf((*MockstockServiceIface)(nil).GetByProduct), stockData)
}

// UpdateStock mocks base method
func (m *MockstockServiceIface) UpdateStock(stockData string) (model.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStock", stockData)
	ret0, _ := ret[0].(model.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStock indicates an expected call of UpdateStock
func (mr *MockstockServiceIfaceMockRecorder) UpdateStock(stockData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStock", reflect.TypeOf((*MockstockServiceIface)(nil).UpdateStock), stockData)
}

// MockunitServiceIface is a mock of unitServiceIface interface
type MockunitServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockunitServiceIfaceMockRecorder
}

// MockunitServiceIfaceMockRecorder is the mock recorder for MockunitServiceIface
type MockunitServiceIfaceMockRecorder struct {
	mock *MockunitServiceIface
}

// NewMockunitServiceIface creates a new mock instance
func NewMockunitServiceIface(ctrl *gomock.Controller) *MockunitServiceIface {
	mock := &MockunitServiceIface{ctrl: ctrl}
	mock.recorder = &MockunitServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockunitServiceIface) EXPECT() *MockunitServiceIfaceMockRecorder {
	return m.recorder
}

// GetAllUnit mocks base method
func (m *MockunitServiceIface) GetAllUnit() ([]model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUnit")
	ret0, _ := ret[0].([]model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUnit indicates an expected call of GetAllUnit
func (mr *MockunitServiceIfaceMockRecorder) GetAllUnit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUnit", reflect.TypeOf((*MockunitServiceIface)(nil).GetAllUnit))
}

// GetOneUnit mocks base method
func (m *MockunitServiceIface) GetOneUnit(id int) (model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneUnit", id)
	ret0, _ := ret[0].(model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneUnit indicates an expected call of GetOneUnit
func (mr *MockunitServiceIfaceMockRecorder) GetOneUnit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneUnit", reflect.TypeOf((*MockunitServiceIface)(nil).GetOneUnit), id)
}

// GetUnitsByName mocks base method
func (m *MockunitServiceIface) GetUnitsByName(name string) ([]model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitsByName", name)
	ret0, _ := ret[0].([]model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitsByName indicates an expected call of GetUnitsByName
func (mr *MockunitServiceIfaceMockRecorder) GetUnitsByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitsByName", reflect.TypeOf((*MockunitServiceIface)(nil).GetUnitsByName), name)
}

// NewUnit mocks base method
func (m *MockunitServiceIface) NewUnit(unitData string) (model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUnit", unitData)
	ret0, _ := ret[0].(model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUnit indicates an expected call of NewUnit
func (mr *MockunitServiceIfaceMockRecorder) NewUnit(unitData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUnit", reflect.TypeOf((*MockunitServiceIface)(nil).NewUnit), unitData)
}

// UpdateUnit mocks base method
func (m *MockunitServiceIface) UpdateUnit(unitData string) (model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnit", unitData)
	ret0, _ := ret[0].(model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUnit indicates an expected call of UpdateUnit
func (mr *MockunitServiceIfaceMockRecorder) UpdateUnit(unitData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnit", reflect.TypeOf((*MockunitServiceIface)(nil).UpdateUnit), unitData)
}

// DeleteUnit mocks base method
func (m *MockunitServiceIface) DeleteUnit(id int) (model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnit", id)
	ret0, _ := ret[0].(model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUnit indicates an expected call of DeleteUnit
func (mr *MockunitServiceIfaceMockRecorder) DeleteUnit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnit", reflect.TypeOf((*MockunitServiceIface)(nil).DeleteUnit), id)
}

// MockuserServiceIface is a mock of userServiceIface interface
type MockuserServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockuserServiceIfaceMockRecorder
}

// MockuserServiceIfaceMockRecorder is the mock recorder for MockuserServiceIface
type MockuserServiceIfaceMockRecorder struct {
	mock *MockuserServiceIface
}

// NewMockuserServiceIface creates a new mock instance
func NewMockuserServiceIface(ctrl *gomock.Controller) *MockuserServiceIface {
	mock := &MockuserServiceIface{ctrl: ctrl}
	mock.recorder = &MockuserServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockuserServiceIface) EXPECT() *MockuserServiceIfaceMockRecorder {
	return m.recorder
}

// GetOneUser mocks base method
func (m *MockuserServiceIface) GetOneUser(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneUser", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneUser indicates an expected call of GetOneUser
func (mr *MockuserServiceIfaceMockRecorder) GetOneUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneUser", reflect.TypeOf((*MockuserServiceIface)(nil).GetOneUser), id)
}

// GetAllUser mocks base method
func (m *MockuserServiceIface) GetAllUser() []model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUser")
	ret0, _ := ret[0].([]model.User)
	return ret0
}

// GetAllUser indicates an expected call of GetAllUser
func (mr *MockuserServiceIfaceMockRecorder) GetAllUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUser", reflect.TypeOf((*MockuserServiceIface)(nil).GetAllUser))
}

// LoginUser mocks base method
func (m *MockuserServiceIface) LoginUser(userData string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", userData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser
func (mr *MockuserServiceIfaceMockRecorder) LoginUser(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockuserServiceIface)(nil).LoginUser), userData)
}

// NewUser mocks base method
func (m *MockuserServiceIface) NewUser(userData string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", userData)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUser indicates an expected call of NewUser
func (mr *MockuserServiceIfaceMockRecorder) NewUser(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockuserServiceIface)(nil).NewUser), userData)
}

// UpdateUser mocks base method
func (m *MockuserServiceIface) UpdateUser(userData string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userData)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockuserServiceIfaceMockRecorder) UpdateUser(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockuserServiceIface)(nil).UpdateUser), userData)
}

// DeleteUser mocks base method
func (m *MockuserServiceIface) DeleteUser(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockuserServiceIfaceMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockuserServiceIface)(nil).DeleteUser), id)
}
