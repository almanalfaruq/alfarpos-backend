package service

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/repository"
)

type OrderService struct {
	Order       repository.IOrderRepository
	OrderDetail repository.IOrderDetailRepository
	Payment     repository.IPaymentRepository
	Customer    repository.ICustomerRepository
	Product     repository.IProductRepository
}

type IOrderService interface {
	GetAllOrder() ([]model.Order, error)
	GetOneOrder(id int) (model.Order, error)
	GetOrderByInvoice(invoice string) (model.Order, error)
	GetOrderByUserId(userId int) ([]model.Order, error)
	NewOrder(OrderData string) (model.Order, error)
	UpdateOrder(OrderData string) (model.Order, error)
	DeleteOrder(id int) (model.Order, error)
}

func (service *OrderService) GetAllOrder() ([]model.Order, error) {
	return service.Order.FindAll(), nil
}

func (service *OrderService) GetOneOrder(id int) (model.Order, error) {
	Order := service.Order.FindById(id)
	if Order.ID == 0 {
		return Order, errors.New("Order not found")
	}
	return Order, nil
}

func (service *OrderService) GetOrderByInvoice(invoice string) (model.Order, error) {
	invoice = strings.ToLower(invoice)
	Order := service.Order.FindByInvoice(invoice)
	if Order.ID == 0 {
		return Order, errors.New("Order not found")
	}
	return Order, nil
}

func (service *OrderService) GetOrderByUserId(userId int) ([]model.Order, error) {
	Orders := service.Order.FindByUserId(userId)
	if len(Orders) == 0 {
		return Orders, errors.New("Orders not found")
	}
	return Orders, nil
}

func (service *OrderService) NewOrder(OrderData string) (model.Order, error) {
	var Order model.Order
	OrderDataByte := []byte(OrderData)
	err := json.Unmarshal(OrderDataByte, &Order)
	if err != nil {
		return Order, err
	}
	Customer := service.Customer.FindById(Order.CustomerID)
	Order.Customer = Customer
	Payment := service.Payment.FindById(Order.PaymentID)
	Order.Payment = Payment
	Order = service.Order.New(Order)
	for _, OrderDetail := range Order.OrderDetails {
		OrderDetail.OrderID = int(Order.ID)
		OrderDetail.Order = Order
		OrderDetail.Product = service.Product.FindById(OrderDetail.ProductID)
		service.OrderDetail.New(OrderDetail)
		Product := OrderDetail.Product
		ProductQty := *Product.Quantity
		stockQty := ProductQty - int64(OrderDetail.Quantity)
		Product.Quantity = &stockQty
		service.Product.Update(Product)
	}
	return Order, nil
}

func (service *OrderService) UpdateOrder(OrderData string) (model.Order, error) {
	var Order model.Order
	OrderDataByte := []byte(OrderData)
	err := json.Unmarshal(OrderDataByte, &Order)
	if err != nil {
		return Order, err
	}
	Order = service.Order.Update(Order)
	return Order, nil
}

func (service *OrderService) DeleteOrder(id int) (model.Order, error) {
	return service.Order.Delete(id)
}
