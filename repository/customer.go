package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/model"
)

type CustomerRepository struct {
	db dbIface
}

func NewCustomerRepo(db dbIface) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (repo *CustomerRepository) FindAll() []model.Customer {
	var customers []model.Customer
	db := repo.db.GetDb()
	db.Find(&customers)
	return customers
}

func (repo *CustomerRepository) FindById(id int64) model.Customer {
	var customer model.Customer
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&customer)
	return customer
}

func (repo *CustomerRepository) New(customer model.Customer) model.Customer {
	db := repo.db.GetDb()
	isNotExists := db.NewRecord(customer)
	if isNotExists {
		db.Create(&customer)
	}
	return customer
}

func (repo *CustomerRepository) Update(customer model.Customer) model.Customer {
	var oldCustomer model.Customer
	db := repo.db.GetDb()
	db.Where("id = ?", customer.ID).First(&oldCustomer)
	oldCustomer = customer
	db.Save(&oldCustomer)
	return customer
}

func (repo *CustomerRepository) Delete(id int64) model.Customer {
	var customer model.Customer
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&customer)
	db.Delete(&customer)
	return customer
}
