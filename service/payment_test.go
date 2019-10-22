package service_test

import (
	"testing"

	"../model"
	. "../service"
	"../test/mocks"
	"../test/resources"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentById(t *testing.T) {
	t.Run("Get Payment By ID - Pass", func(t *testing.T) {
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("FindById", 1).Return(resources.Payment1)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := resources.Payment1

		actualResult, err := paymentService.GetOnePayment(1)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Get Payment By ID - Error", func(t *testing.T) {
		var payment model.Payment
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("FindById", 5).Return(payment)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := payment

		actualResult, err := paymentService.GetOnePayment(5)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestGetAllPayment(t *testing.T) {
	paymentRepository := new(mocks.PaymentRepository)

	paymentRepository.On("FindAll").Return(resources.Payments)

	paymentService := PaymentService{
		IPaymentRepository: paymentRepository,
	}

	expectedResult := resources.Payments

	actualResult := paymentService.GetAllPayment()

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

func TestNewPayment(t *testing.T) {
	t.Run("New Payment - Success", func(t *testing.T) {
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("New", resources.Payment2).Return(resources.Payment2)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := resources.Payment2

		jsonPayment := `{"id": 2, "name": "Payment2"}`

		actualResult, err := paymentService.NewPayment(jsonPayment)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("New Payment - Error", func(t *testing.T) {
		var payment model.Payment
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("New", payment).Return(payment)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := payment

		jsonPayment := `{name: "Payment2"}`

		actualResult, err := paymentService.NewPayment(jsonPayment)

		assert.NotNil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestUpdatePayment(t *testing.T) {
	t.Run("Update Payment - Success", func(t *testing.T) {
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("Update", resources.Payment2Updated).Return(resources.Payment2Updated)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := resources.Payment2Updated

		jsonPayment := `{"id": 2, "name": "Payment2Updated"}`

		actualResult, err := paymentService.UpdatePayment(jsonPayment)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Update Payment - Error", func(t *testing.T) {
		var payment model.Payment
		paymentRepository := new(mocks.PaymentRepository)

		paymentRepository.On("Update", payment).Return(payment)

		paymentService := PaymentService{
			IPaymentRepository: paymentRepository,
		}

		expectedResult := payment

		jsonPayment := `{name: "Payment2Updated"}`

		actualResult, err := paymentService.UpdatePayment(jsonPayment)

		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestDeletePayment(t *testing.T) {
	paymentRepository := new(mocks.PaymentRepository)

	paymentRepository.On("Delete", 4).Return(resources.Payment4)

	paymentService := PaymentService{
		IPaymentRepository: paymentRepository,
	}

	expectedResult := resources.Payment4

	id := 4

	actualResult, err := paymentService.DeletePayment(id)

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult.ID, actualResult.ID)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
}
