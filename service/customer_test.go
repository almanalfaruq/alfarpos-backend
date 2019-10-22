package service_test

import (
	"fmt"
	"testing"
	"time"

	"../model"
	. "../service"
	"../test/mocks"
	"../test/resources"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerById(t *testing.T) {
	t.Run("Get Customer By ID - Pass", func(t *testing.T) {
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("FindById", 1).Return(resources.Customer1)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := resources.Customer1

		actualResult, err := customerService.GetOneCustomer(1)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Code, actualResult.Code)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
		assert.Equal(t, expectedResult.Address, actualResult.Address)
		assert.Equal(t, expectedResult.Phone, actualResult.Phone)
	})

	t.Run("Get Customer By ID - Error", func(t *testing.T) {
		var customer model.Customer
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("FindById", 5).Return(customer)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := customer

		actualResult, err := customerService.GetOneCustomer(5)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestGetAllCustomer(t *testing.T) {
	customerRepository := new(mocks.CustomerRepository)

	customerRepository.On("FindAll").Return(resources.Customers)

	customerService := CustomerService{
		ICustomerRepository: customerRepository,
	}

	expectedResult := resources.Customers

	actualResult := customerService.GetAllCustomer()

	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
	assert.Equal(t, expectedResult[2].ID, actualResult[2].ID)
	assert.Equal(t, expectedResult[2].Name, actualResult[2].Name)
	assert.Equal(t, expectedResult[3].ID, actualResult[3].ID)
	assert.Equal(t, expectedResult[3].Name, actualResult[3].Name)
}

func TestNewCustomer(t *testing.T) {
	t.Run("New Customer - Success", func(t *testing.T) {
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("New", resources.Customer2).Return(resources.Customer2)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := resources.Customer2

		jsonCustomer := `{
			"id": 2,
			"name": "Customer2",
			"address": "Boyolali",
			"phone": "081225812599"
		}`

		actualResult, err := customerService.NewCustomer(jsonCustomer)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Code, actualResult.Code)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
		assert.Equal(t, expectedResult.Address, actualResult.Address)
		assert.Equal(t, expectedResult.Phone, actualResult.Phone)
	})

	t.Run("New Customer - Error", func(t *testing.T) {
		var customer model.Customer
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("New", customer).Return(customer)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := customer

		jsonCustomer := `{name: "Customer2"}`

		actualResult, err := customerService.NewCustomer(jsonCustomer)

		assert.NotNil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestUpdateCustomer(t *testing.T) {
	t.Run("Update Customer - Success", func(t *testing.T) {
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("Update", resources.Customer2Updated).Return(resources.Customer2Updated)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := resources.Customer2Updated
		now := time.Now().Format("20060102150405")

		jsonCustomer := fmt.Sprintf(`{
			"id": 2,
			"code": "CUST-%s",
			"name": "Customer2Updated",
			"address": "Boyolali",
			"phone": "081225812599"
		}`, now)

		actualResult, err := customerService.UpdateCustomer(jsonCustomer)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Code, actualResult.Code)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
		assert.Equal(t, expectedResult.Address, actualResult.Address)
		assert.Equal(t, expectedResult.Phone, actualResult.Phone)
	})

	t.Run("Update Customer - Error", func(t *testing.T) {
		var customer model.Customer
		customerRepository := new(mocks.CustomerRepository)

		customerRepository.On("Update", customer).Return(customer)

		customerService := CustomerService{
			ICustomerRepository: customerRepository,
		}

		expectedResult := customer

		jsonCustomer := `{name: "Customer2Updated"}`

		actualResult, err := customerService.UpdateCustomer(jsonCustomer)

		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestDeleteCustomer(t *testing.T) {
	customerRepository := new(mocks.CustomerRepository)

	customerRepository.On("Delete", 4).Return(resources.Customer4)

	customerService := CustomerService{
		ICustomerRepository: customerRepository,
	}

	expectedResult := resources.Customer4

	id := 4

	actualResult, err := customerService.DeleteCustomer(id)

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult.ID, actualResult.ID)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
}
