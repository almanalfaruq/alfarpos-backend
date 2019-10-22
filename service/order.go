package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type OrderService struct {
	order    repository.IOrderRepository
	payment  repository.IPaymentRepository
	customer repository.ICustomerRepository
}

type IOrderService interface {
	GetAllOrder() []model.Order
	GetOneOrder(id int) (model.Order, error)
	GetOrderByInvoice(orderData string) (model.Order, error)
	GetOrderByUserId(orderData string) ([]model.Order, error)
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

func (service *OrderService) GetOrderByInvoice(orderData string) (model.Order, error) {
	var order model.Order
	orderDataByte := []byte(orderData)
	err := json.Unmarshal(orderDataByte, &order)
	if err != nil {
		return model.Order{}, err
	}
	return service.order.FindByInvoice(order.Invoice), nil
}

func (service *OrderService) GetOrderByUserId(orderData string) ([]model.Order, error) {
	var order model.Order
	orderDataByte := []byte(orderData)
	err := json.Unmarshal(orderDataByte, &order)
	if err != nil {
		return []model.Order{}, err
	}
	return service.order.FindByUserId(order.UserID), nil
}

func (service *OrderService) NewOrder(orderData string) (model.Order, error) {
	var order model.Order
	orderDataByte := []byte(orderData)
	err := json.Unmarshal(orderDataByte, &order)
	if err != nil {
		return order, err
	}
	return service.order.New(order), nil
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
