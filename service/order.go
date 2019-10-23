package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type OrderService struct {
	order       repository.IOrderRepository
	orderDetail repository.IOrderDetailRepository
	payment     repository.IPaymentRepository
	customer    repository.ICustomerRepository
}

type IOrderService interface {
	GetAllOrder() []model.Order
	GetOneOrder(id int) (model.Order, error)
	GetOrderByInvoice(invoice string) (model.Order, error)
	GetOrderByUserId(userId int) ([]model.Order, error)
	NewOrder(orderData string) (model.Order, error)
	UpdateOrder(orderData string) (model.Order, error)
	DeleteOrder(id int) int
}

func (service *OrderService) GetAllOrder() []model.Order {
	return service.order.FindAll()
}

func (service *OrderService) GetOneOrder(id int) (model.Order, error) {
	return service.order.FindById(id), nil
}

func (service *OrderService) GetOrderByInvoice(invoice string) (model.Order, error) {
	return service.order.FindByInvoice(invoice), nil
}

func (service *OrderService) GetOrderByUserId(userId int) ([]model.Order, error) {
	return service.order.FindByUserId(userId), nil
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
		service.orderDetail.New(orderDetail)
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
	return service.order.Delete(id), nil
}
