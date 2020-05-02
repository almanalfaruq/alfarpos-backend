// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package service is a generated GoMock package.
package service

import (
	model "github.com/almanalfaruq/alfarpos-backend/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockcategoryRepositoryIface is a mock of categoryRepositoryIface interface
type MockcategoryRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockcategoryRepositoryIfaceMockRecorder
}

// MockcategoryRepositoryIfaceMockRecorder is the mock recorder for MockcategoryRepositoryIface
type MockcategoryRepositoryIfaceMockRecorder struct {
	mock *MockcategoryRepositoryIface
}

// NewMockcategoryRepositoryIface creates a new mock instance
func NewMockcategoryRepositoryIface(ctrl *gomock.Controller) *MockcategoryRepositoryIface {
	mock := &MockcategoryRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockcategoryRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcategoryRepositoryIface) EXPECT() *MockcategoryRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockcategoryRepositoryIface) FindAll() []model.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Category)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockcategoryRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockcategoryRepositoryIface) FindById(id int64) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockcategoryRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).FindById), id)
}

// FindByName mocks base method
func (m *MockcategoryRepositoryIface) FindByName(name string) []model.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].([]model.Category)
	return ret0
}

// FindByName indicates an expected call of FindByName
func (mr *MockcategoryRepositoryIfaceMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).FindByName), name)
}

// New mocks base method
func (m *MockcategoryRepositoryIface) New(category model.Category) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", category)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockcategoryRepositoryIfaceMockRecorder) New(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).New), category)
}

// Update mocks base method
func (m *MockcategoryRepositoryIface) Update(category model.Category) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", category)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockcategoryRepositoryIfaceMockRecorder) Update(category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).Update), category)
}

// Delete mocks base method
func (m *MockcategoryRepositoryIface) Delete(id int64) (model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockcategoryRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcategoryRepositoryIface)(nil).Delete), id)
}

// MockcustomerRepositoryIface is a mock of customerRepositoryIface interface
type MockcustomerRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockcustomerRepositoryIfaceMockRecorder
}

// MockcustomerRepositoryIfaceMockRecorder is the mock recorder for MockcustomerRepositoryIface
type MockcustomerRepositoryIfaceMockRecorder struct {
	mock *MockcustomerRepositoryIface
}

// NewMockcustomerRepositoryIface creates a new mock instance
func NewMockcustomerRepositoryIface(ctrl *gomock.Controller) *MockcustomerRepositoryIface {
	mock := &MockcustomerRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockcustomerRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockcustomerRepositoryIface) EXPECT() *MockcustomerRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockcustomerRepositoryIface) FindAll() []model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Customer)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockcustomerRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockcustomerRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockcustomerRepositoryIface) FindById(id int64) model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Customer)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockcustomerRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockcustomerRepositoryIface)(nil).FindById), id)
}

// New mocks base method
func (m *MockcustomerRepositoryIface) New(customer model.Customer) model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", customer)
	ret0, _ := ret[0].(model.Customer)
	return ret0
}

// New indicates an expected call of New
func (mr *MockcustomerRepositoryIfaceMockRecorder) New(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockcustomerRepositoryIface)(nil).New), customer)
}

// Update mocks base method
func (m *MockcustomerRepositoryIface) Update(customer model.Customer) model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", customer)
	ret0, _ := ret[0].(model.Customer)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockcustomerRepositoryIfaceMockRecorder) Update(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockcustomerRepositoryIface)(nil).Update), customer)
}

// Delete mocks base method
func (m *MockcustomerRepositoryIface) Delete(id int64) model.Customer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Customer)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockcustomerRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcustomerRepositoryIface)(nil).Delete), id)
}

// MockorderDetailRepositoryIface is a mock of orderDetailRepositoryIface interface
type MockorderDetailRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockorderDetailRepositoryIfaceMockRecorder
}

// MockorderDetailRepositoryIfaceMockRecorder is the mock recorder for MockorderDetailRepositoryIface
type MockorderDetailRepositoryIfaceMockRecorder struct {
	mock *MockorderDetailRepositoryIface
}

// NewMockorderDetailRepositoryIface creates a new mock instance
func NewMockorderDetailRepositoryIface(ctrl *gomock.Controller) *MockorderDetailRepositoryIface {
	mock := &MockorderDetailRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockorderDetailRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockorderDetailRepositoryIface) EXPECT() *MockorderDetailRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindByOrder mocks base method
func (m *MockorderDetailRepositoryIface) FindByOrder(order model.Order) ([]model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByOrder", order)
	ret0, _ := ret[0].([]model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByOrder indicates an expected call of FindByOrder
func (mr *MockorderDetailRepositoryIfaceMockRecorder) FindByOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByOrder", reflect.TypeOf((*MockorderDetailRepositoryIface)(nil).FindByOrder), order)
}

// New mocks base method
func (m *MockorderDetailRepositoryIface) New(orderDetail model.OrderDetail) (model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", orderDetail)
	ret0, _ := ret[0].(model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockorderDetailRepositoryIfaceMockRecorder) New(orderDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockorderDetailRepositoryIface)(nil).New), orderDetail)
}

// Update mocks base method
func (m *MockorderDetailRepositoryIface) Update(orderDetail model.OrderDetail) (model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", orderDetail)
	ret0, _ := ret[0].(model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockorderDetailRepositoryIfaceMockRecorder) Update(orderDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockorderDetailRepositoryIface)(nil).Update), orderDetail)
}

// Delete mocks base method
func (m *MockorderDetailRepositoryIface) Delete(id int64) (model.OrderDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.OrderDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockorderDetailRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockorderDetailRepositoryIface)(nil).Delete), id)
}

// DeleteByOrderId mocks base method
func (m *MockorderDetailRepositoryIface) DeleteByOrderId(id int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByOrderId", id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByOrderId indicates an expected call of DeleteByOrderId
func (mr *MockorderDetailRepositoryIfaceMockRecorder) DeleteByOrderId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByOrderId", reflect.TypeOf((*MockorderDetailRepositoryIface)(nil).DeleteByOrderId), id)
}

// MockorderRepositoryIface is a mock of orderRepositoryIface interface
type MockorderRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockorderRepositoryIfaceMockRecorder
}

// MockorderRepositoryIfaceMockRecorder is the mock recorder for MockorderRepositoryIface
type MockorderRepositoryIfaceMockRecorder struct {
	mock *MockorderRepositoryIface
}

// NewMockorderRepositoryIface creates a new mock instance
func NewMockorderRepositoryIface(ctrl *gomock.Controller) *MockorderRepositoryIface {
	mock := &MockorderRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockorderRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockorderRepositoryIface) EXPECT() *MockorderRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockorderRepositoryIface) FindAll() []model.Order {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Order)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockorderRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockorderRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockorderRepositoryIface) FindById(id int64) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockorderRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockorderRepositoryIface)(nil).FindById), id)
}

// FindByInvoice mocks base method
func (m *MockorderRepositoryIface) FindByInvoice(invoice string) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByInvoice", invoice)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByInvoice indicates an expected call of FindByInvoice
func (mr *MockorderRepositoryIfaceMockRecorder) FindByInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByInvoice", reflect.TypeOf((*MockorderRepositoryIface)(nil).FindByInvoice), invoice)
}

// FindByUserId mocks base method
func (m *MockorderRepositoryIface) FindByUserId(userId int64) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", userId)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserId indicates an expected call of FindByUserId
func (mr *MockorderRepositoryIfaceMockRecorder) FindByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockorderRepositoryIface)(nil).FindByUserId), userId)
}

// New mocks base method
func (m *MockorderRepositoryIface) New(order model.Order) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", order)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockorderRepositoryIfaceMockRecorder) New(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockorderRepositoryIface)(nil).New), order)
}

// Update mocks base method
func (m *MockorderRepositoryIface) Update(order model.Order) model.Order {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", order)
	ret0, _ := ret[0].(model.Order)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockorderRepositoryIfaceMockRecorder) Update(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockorderRepositoryIface)(nil).Update), order)
}

// Delete mocks base method
func (m *MockorderRepositoryIface) Delete(id int64) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockorderRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockorderRepositoryIface)(nil).Delete), id)
}

// MockpaymentRepositoryIface is a mock of paymentRepositoryIface interface
type MockpaymentRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockpaymentRepositoryIfaceMockRecorder
}

// MockpaymentRepositoryIfaceMockRecorder is the mock recorder for MockpaymentRepositoryIface
type MockpaymentRepositoryIfaceMockRecorder struct {
	mock *MockpaymentRepositoryIface
}

// NewMockpaymentRepositoryIface creates a new mock instance
func NewMockpaymentRepositoryIface(ctrl *gomock.Controller) *MockpaymentRepositoryIface {
	mock := &MockpaymentRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockpaymentRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockpaymentRepositoryIface) EXPECT() *MockpaymentRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockpaymentRepositoryIface) FindAll() []model.Payment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Payment)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockpaymentRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockpaymentRepositoryIface) FindById(id int64) model.Payment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Payment)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockpaymentRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).FindById), id)
}

// FindByName mocks base method
func (m *MockpaymentRepositoryIface) FindByName(name string) []model.Payment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].([]model.Payment)
	return ret0
}

// FindByName indicates an expected call of FindByName
func (mr *MockpaymentRepositoryIfaceMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).FindByName), name)
}

// New mocks base method
func (m *MockpaymentRepositoryIface) New(payment model.Payment) model.Payment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", payment)
	ret0, _ := ret[0].(model.Payment)
	return ret0
}

// New indicates an expected call of New
func (mr *MockpaymentRepositoryIfaceMockRecorder) New(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).New), payment)
}

// Update mocks base method
func (m *MockpaymentRepositoryIface) Update(payment model.Payment) model.Payment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", payment)
	ret0, _ := ret[0].(model.Payment)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockpaymentRepositoryIfaceMockRecorder) Update(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).Update), payment)
}

// Delete mocks base method
func (m *MockpaymentRepositoryIface) Delete(id int64) (model.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockpaymentRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockpaymentRepositoryIface)(nil).Delete), id)
}

// MockproductRepositoryIface is a mock of productRepositoryIface interface
type MockproductRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockproductRepositoryIfaceMockRecorder
}

// MockproductRepositoryIfaceMockRecorder is the mock recorder for MockproductRepositoryIface
type MockproductRepositoryIfaceMockRecorder struct {
	mock *MockproductRepositoryIface
}

// NewMockproductRepositoryIface creates a new mock instance
func NewMockproductRepositoryIface(ctrl *gomock.Controller) *MockproductRepositoryIface {
	mock := &MockproductRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockproductRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockproductRepositoryIface) EXPECT() *MockproductRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockproductRepositoryIface) FindAll() []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockproductRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockproductRepositoryIface) FindById(id int64) model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Product)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockproductRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindById), id)
}

// FindByCode mocks base method
func (m *MockproductRepositoryIface) FindByCode(code string) []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCode", code)
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// FindByCode indicates an expected call of FindByCode
func (mr *MockproductRepositoryIfaceMockRecorder) FindByCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCode", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindByCode), code)
}

// FindByName mocks base method
func (m *MockproductRepositoryIface) FindByName(name string) []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// FindByName indicates an expected call of FindByName
func (mr *MockproductRepositoryIfaceMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindByName), name)
}

// FindByCategoryName mocks base method
func (m *MockproductRepositoryIface) FindByCategoryName(name string) []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCategoryName", name)
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// FindByCategoryName indicates an expected call of FindByCategoryName
func (mr *MockproductRepositoryIfaceMockRecorder) FindByCategoryName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCategoryName", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindByCategoryName), name)
}

// FindByUnitName mocks base method
func (m *MockproductRepositoryIface) FindByUnitName(name string) []model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUnitName", name)
	ret0, _ := ret[0].([]model.Product)
	return ret0
}

// FindByUnitName indicates an expected call of FindByUnitName
func (mr *MockproductRepositoryIfaceMockRecorder) FindByUnitName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUnitName", reflect.TypeOf((*MockproductRepositoryIface)(nil).FindByUnitName), name)
}

// New mocks base method
func (m *MockproductRepositoryIface) New(product model.Product) model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", product)
	ret0, _ := ret[0].(model.Product)
	return ret0
}

// New indicates an expected call of New
func (mr *MockproductRepositoryIfaceMockRecorder) New(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockproductRepositoryIface)(nil).New), product)
}

// Update mocks base method
func (m *MockproductRepositoryIface) Update(product model.Product) model.Product {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", product)
	ret0, _ := ret[0].(model.Product)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockproductRepositoryIfaceMockRecorder) Update(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockproductRepositoryIface)(nil).Update), product)
}

// Delete mocks base method
func (m *MockproductRepositoryIface) Delete(id int64) (model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockproductRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockproductRepositoryIface)(nil).Delete), id)
}

// DeleteAll mocks base method
func (m *MockproductRepositoryIface) DeleteAll() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(int64)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockproductRepositoryIfaceMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockproductRepositoryIface)(nil).DeleteAll))
}

// MockstockRepositoryIface is a mock of stockRepositoryIface interface
type MockstockRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockstockRepositoryIfaceMockRecorder
}

// MockstockRepositoryIfaceMockRecorder is the mock recorder for MockstockRepositoryIface
type MockstockRepositoryIfaceMockRecorder struct {
	mock *MockstockRepositoryIface
}

// NewMockstockRepositoryIface creates a new mock instance
func NewMockstockRepositoryIface(ctrl *gomock.Controller) *MockstockRepositoryIface {
	mock := &MockstockRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockstockRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstockRepositoryIface) EXPECT() *MockstockRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockstockRepositoryIface) FindAll() []model.Stock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Stock)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockstockRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockstockRepositoryIface)(nil).FindAll))
}

// FindByProduct mocks base method
func (m *MockstockRepositoryIface) FindByProduct(product model.Product) model.Stock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByProduct", product)
	ret0, _ := ret[0].(model.Stock)
	return ret0
}

// FindByProduct indicates an expected call of FindByProduct
func (mr *MockstockRepositoryIfaceMockRecorder) FindByProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByProduct", reflect.TypeOf((*MockstockRepositoryIface)(nil).FindByProduct), product)
}

// New mocks base method
func (m *MockstockRepositoryIface) New(stock model.Stock) model.Stock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", stock)
	ret0, _ := ret[0].(model.Stock)
	return ret0
}

// New indicates an expected call of New
func (mr *MockstockRepositoryIfaceMockRecorder) New(stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockstockRepositoryIface)(nil).New), stock)
}

// Update mocks base method
func (m *MockstockRepositoryIface) Update(stock model.Stock) model.Stock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", stock)
	ret0, _ := ret[0].(model.Stock)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockstockRepositoryIfaceMockRecorder) Update(stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockstockRepositoryIface)(nil).Update), stock)
}

// Delete mocks base method
func (m *MockstockRepositoryIface) Delete(id int64) (model.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockstockRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockstockRepositoryIface)(nil).Delete), id)
}

// DeleteAll mocks base method
func (m *MockstockRepositoryIface) DeleteAll() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(int64)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll
func (mr *MockstockRepositoryIfaceMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockstockRepositoryIface)(nil).DeleteAll))
}

// MockunitRepositoryIface is a mock of unitRepositoryIface interface
type MockunitRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockunitRepositoryIfaceMockRecorder
}

// MockunitRepositoryIfaceMockRecorder is the mock recorder for MockunitRepositoryIface
type MockunitRepositoryIfaceMockRecorder struct {
	mock *MockunitRepositoryIface
}

// NewMockunitRepositoryIface creates a new mock instance
func NewMockunitRepositoryIface(ctrl *gomock.Controller) *MockunitRepositoryIface {
	mock := &MockunitRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockunitRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockunitRepositoryIface) EXPECT() *MockunitRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockunitRepositoryIface) FindAll() []model.Unit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.Unit)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockunitRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockunitRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockunitRepositoryIface) FindById(id int64) model.Unit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.Unit)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockunitRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockunitRepositoryIface)(nil).FindById), id)
}

// FindByName mocks base method
func (m *MockunitRepositoryIface) FindByName(name string) []model.Unit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].([]model.Unit)
	return ret0
}

// FindByName indicates an expected call of FindByName
func (mr *MockunitRepositoryIfaceMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockunitRepositoryIface)(nil).FindByName), name)
}

// New mocks base method
func (m *MockunitRepositoryIface) New(unit model.Unit) model.Unit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", unit)
	ret0, _ := ret[0].(model.Unit)
	return ret0
}

// New indicates an expected call of New
func (mr *MockunitRepositoryIfaceMockRecorder) New(unit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockunitRepositoryIface)(nil).New), unit)
}

// Update mocks base method
func (m *MockunitRepositoryIface) Update(unit model.Unit) model.Unit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", unit)
	ret0, _ := ret[0].(model.Unit)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockunitRepositoryIfaceMockRecorder) Update(unit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockunitRepositoryIface)(nil).Update), unit)
}

// Delete mocks base method
func (m *MockunitRepositoryIface) Delete(id int64) (model.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockunitRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockunitRepositoryIface)(nil).Delete), id)
}

// MockuserRepositoryIface is a mock of userRepositoryIface interface
type MockuserRepositoryIface struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepositoryIfaceMockRecorder
}

// MockuserRepositoryIfaceMockRecorder is the mock recorder for MockuserRepositoryIface
type MockuserRepositoryIfaceMockRecorder struct {
	mock *MockuserRepositoryIface
}

// NewMockuserRepositoryIface creates a new mock instance
func NewMockuserRepositoryIface(ctrl *gomock.Controller) *MockuserRepositoryIface {
	mock := &MockuserRepositoryIface{ctrl: ctrl}
	mock.recorder = &MockuserRepositoryIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockuserRepositoryIface) EXPECT() *MockuserRepositoryIfaceMockRecorder {
	return m.recorder
}

// FindAll mocks base method
func (m *MockuserRepositoryIface) FindAll() []model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]model.User)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockuserRepositoryIfaceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockuserRepositoryIface)(nil).FindAll))
}

// FindById mocks base method
func (m *MockuserRepositoryIface) FindById(id int64) model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(model.User)
	return ret0
}

// FindById indicates an expected call of FindById
func (mr *MockuserRepositoryIfaceMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockuserRepositoryIface)(nil).FindById), id)
}

// FindByUsername mocks base method
func (m *MockuserRepositoryIface) FindByUsername(username string) model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", username)
	ret0, _ := ret[0].(model.User)
	return ret0
}

// FindByUsername indicates an expected call of FindByUsername
func (mr *MockuserRepositoryIfaceMockRecorder) FindByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockuserRepositoryIface)(nil).FindByUsername), username)
}

// New mocks base method
func (m *MockuserRepositoryIface) New(user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New
func (mr *MockuserRepositoryIfaceMockRecorder) New(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockuserRepositoryIface)(nil).New), user)
}

// Update mocks base method
func (m *MockuserRepositoryIface) Update(user model.User) model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(model.User)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockuserRepositoryIfaceMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockuserRepositoryIface)(nil).Update), user)
}

// Delete mocks base method
func (m *MockuserRepositoryIface) Delete(id int64) model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(model.User)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockuserRepositoryIfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockuserRepositoryIface)(nil).Delete), id)
}
