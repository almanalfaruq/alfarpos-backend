package repository

import (
	"../model"
	"../util"
)

type PaymentRepository struct {
	util.IDatabaseConnection
}

type IPaymentRepository interface {
	FindAll() []model.Payment
	FindById(id int) model.Payment
	New(payment model.Payment) model.Payment
	Update(payment model.Payment) model.Payment
	Delete(id int) model.Payment
}

func (repo *PaymentRepository) FindAll() []model.Payment {
	var categories []model.Payment
	db := repo.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *PaymentRepository) FindById(id int) model.Payment {
	var payment model.Payment
	db := repo.GetDb()
	db.Where("id = ?", id).First(&payment)
	return payment
}

func (repo *PaymentRepository) New(payment model.Payment) model.Payment {
	db := repo.GetDb()
	isNotExist := db.NewRecord(payment)
	if isNotExist {
		db.Create(&payment)
	}
	return payment
}

func (repo *PaymentRepository) Update(payment model.Payment) model.Payment {
	var oldPayment model.Payment
	db := repo.GetDb()
	db.Where("id = ?", payment.ID).First(&oldPayment)
	oldPayment = payment
	db.Save(&oldPayment)
	return payment
}

func (repo *PaymentRepository) Delete(id int) model.Payment {
	var payment model.Payment
	db := repo.GetDb()
	db.Where("id = ?", id).First(&payment)
	db.Delete(&payment)
	return payment
}
