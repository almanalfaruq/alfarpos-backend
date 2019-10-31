package service

import (
	"encoding/json"
	"errors"
	"strings"

	"../model"
	"../repository"
)

type OrderService struct {
	order       repository.IOrderRepository
	orderDetail repository.IOrderDetailRepository
	payment     repository.IPaymentRepository
	customer    repository.ICustomerRepository
	product     repository.IProductRepository
}

type IOrderService interface {
	GetAllOrder() ([]model.Order, error)
	GetOneOrder(id int) (model.Order, error)
	GetOrderByInvoice(invoice string) (model.Order, error)
	GetOrderByUserId(userId int) ([]model.Order, error)
	NewOrder(orderData string) (model.Order, error)
	UpdateOrder(orderData string) (model.Order, error)
	DeleteOrder(id int) int
}

func (service *OrderService) GetAllOrder() ([]model.Order, error) {
	return service.order.FindAll(), nil
}

func (service *OrderService) GetOneOrder(id int) (model.Order, error) {
	order := service.order.FindById(id)
	if order.ID == 0 {
		return order, errors.New("Order not found")
	}
	return order, nil
}

func (service *OrderService) GetOrderByInvoice(invoice string) (model.Order, error) {
	invoice = strings.ToLower(invoice)
	order := service.order.FindByInvoice(invoice)
	if order.ID == 0 {
		return order, errors.New("Order not found")
	}
	return order, nil
}

func (service *OrderService) GetOrderByUserId(userId int) ([]model.Order, error) {
	orders := service.order.FindByUserId(userId)
	if len(orders) == 0 {
		return orders, errors.New("Orders not found")
	}
	return orders, nil
}

func (service *OrderService) NewOrder(orderData string) (model.Order, error) {
	var order model.Order
	orderDataByte := []byte(orderData)
	err := json.Unmarshal(orderDataByte, &order)
	if err != nil {
		return order, err
	}
	customer := service.customer.FindById(order.CustomerID)
	order.Customer = customer
	payment := service.payment.FindById(order.PaymentID)
	order.Payment = payment
	order = service.order.New(order)
	for _, orderDetail := range order.OrderDetails {
		orderDetail.OrderID = int(order.ID)
		orderDetail.Order = order
		orderDetail.Product = service.product.FindById(orderDetail.ProductID)
		service.orderDetail.New(orderDetail)
		product := orderDetail.Product
		productQty := *product.Quantity
		stockQty := productQty - int64(orderDetail.Quantity)
		product.Quantity = &stockQty
		service.product.Update(product)
	}
	return order, nil
}

func (service *OrderService) UpdateOrder(orderData string) (model.Order, error) {
	var order model.Order
	orderDataByte := []byte(orderData)
	err := json.Unmarshal(orderDataByte, &order)
	if err != nil {
		return order, err
	}
	order = service.order.Update(order)
	return order, nil
}

func (service *OrderService) DeleteOrder(id int) (model.Order, error) {
	return service.order.Delete(id)
}
