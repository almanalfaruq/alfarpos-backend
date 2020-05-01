package service

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type PaymentService struct {
	payment paymentRepositoryIface
}

func NewPaymentService(paymentRepo paymentRepositoryIface) *PaymentService {
	return &PaymentService{
		payment: paymentRepo,
	}
}

func (service *PaymentService) GetAllPayment() ([]model.Payment, error) {
	return service.payment.FindAll(), nil
}

func (service *PaymentService) GetOnePayment(id int) (model.Payment, error) {
	payment := service.payment.FindById(id)
	if payment.ID == 0 {
		return payment, errors.New("Payment not found")
	}
	return payment, nil
}

func (service *PaymentService) GetPaymentsByName(name string) ([]model.Payment, error) {
	name = strings.ToLower(name)
	payments := service.payment.FindByName(name)
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
	return service.payment.New(payment), nil
}

func (service *PaymentService) UpdatePayment(paymentData string) (model.Payment, error) {
	var payment model.Payment
	paymentDataByte := []byte(paymentData)
	err := json.Unmarshal(paymentDataByte, &payment)
	if err != nil {
		return payment, err
	}
	payment = service.payment.Update(payment)
	return payment, nil
}

func (service *PaymentService) DeletePayment(id int) (model.Payment, error) {
	return service.payment.Delete(id)
}
