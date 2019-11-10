package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/repository"
)

type CustomerService struct {
	repository.ICustomerRepository
}

type ICustomerService interface {
	GetOneCustomer(id int) (model.Customer, error)
	GetAllCustomer() []model.Customer
	NewCustomer(customerData string) (model.Customer, error)
	UpdateCustomer(customerData string) (model.Customer, error)
	DeleteCustomer(id int) (model.Customer, error)
}

func (service *CustomerService) GetAllCustomer() []model.Customer {
	return service.FindAll()
}

func (service *CustomerService) GetOneCustomer(id int) (model.Customer, error) {
	return service.FindById(id), nil
}

func (service *CustomerService) NewCustomer(customerData string) (model.Customer, error) {
	var customer model.Customer
	customerDataByte := []byte(customerData)
	err := json.Unmarshal(customerDataByte, &customer)
	if err != nil {
		return customer, err
	}
	customer.Code = fmt.Sprintf("CUST-%s", time.Now().Format("20060102150405"))
	return service.New(customer), nil
}

func (service *CustomerService) UpdateCustomer(customerData string) (model.Customer, error) {
	var customer model.Customer
	customerDataByte := []byte(customerData)
	err := json.Unmarshal(customerDataByte, &customer)
	if err != nil {
		return customer, err
	}
	customer = service.Update(customer)
	return customer, nil
}

func (service *CustomerService) DeleteCustomer(id int) (model.Customer, error) {
	return service.Delete(id), nil
}
