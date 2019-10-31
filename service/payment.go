package service

import (
	"encoding/json"
	"errors"
	"strings"

	"../model"
	"../repository"
)

type PaymentService struct {
	repository.IPaymentRepository
}

type IPaymentService interface {
	GetAllPayment() ([]model.Payment, error)
	GetOnePayment(id int) (model.Payment, error)
	GetPaymentsByName(name string) ([]model.Payment, error)
	NewPayment(paymentData string) (model.Payment, error)
	UpdatePayment(paymentData string) (model.Payment, error)
	DeletePayment(id int) (model.Payment, error)
}

func (service *PaymentService) GetAllPayment() ([]model.Payment, error) {
	return service.FindAll(), nil
}

func (service *PaymentService) GetOnePayment(id int) (model.Payment, error) {
	payment := service.FindById(id)
	if payment.ID == 0 {
		return payment, errors.New("Payment not found")
	}
	return payment, nil
}

func (service *PaymentService) GetPaymentsByName(name string) ([]model.Payment, error) {
	name = strings.ToLower(name)
	payments := service.FindByName(name)
	if len(payments) == 0 {
		return payments, errors.New("Payments not found")
	}
	return payments, nil
}

func (service *PaymentService) NewPayment(paymentData string) (model.Payment, error) {
	var payment model.Payment
	paymentDataByte := []byte(paymentData)
	err := json.Unmarshal(paymentDataByte, &payment)
	if err != nil {
		return payment, err
	}
	return service.New(payment), nil
}

func (service *PaymentService) UpdatePayment(paymentData string) (model.Payment, error) {
	var payment model.Payment
	paymentDataByte := []byte(paymentData)
	err := json.Unmarshal(paymentDataByte, &payment)
	if err != nil {
		return payment, err
	}
	payment = service.Update(payment)
	return payment, nil
}

func (service *PaymentService) DeletePayment(id int) (model.Payment, error) {
	return service.Delete(id)
}
