package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type PaymentService struct {
	repository.IPaymentRepository
}

type IPaymentService interface {
	GetAllPayment() []model.Payment
	GetOnePayment(id int) (model.Payment, error)
	NewPayment(paymentData string) (model.Payment, error)
	UpdatePayment(paymentData string) (model.Payment, error)
	DeletePayment(id int) int
}

func (service *PaymentService) GetAllPayment() []model.Payment {
	return service.FindAll()
}

func (service *PaymentService) GetOnePayment(id int) (model.Payment, error) {
	return service.FindById(id), nil
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
	return service.Delete(id), nil
}
