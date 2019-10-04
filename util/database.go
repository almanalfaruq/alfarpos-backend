package util

import (
	"fmt"

	"../model"
	"github.com/jinzhu/gorm"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

type IDatabaseConnection interface {
	Open(config Config) *gorm.DB
	GetDb() *gorm.DB
	MigrateDb()
	DropDb()
}

func (databaseConnection *DatabaseConnection) Open(config Config) *gorm.DB {
	var url string
	if config.Env == "test" {
		url = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBTestName)
	} else {
		url = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName)
	}
	var err error
	databaseConnection.DB, err = gorm.Open("postgres", url)
	if err != nil {
		panic("Cannot connect to the database")
	}
	return databaseConnection.DB
}

func (databaseConnection *DatabaseConnection) GetDb() *gorm.DB {
	return databaseConnection.DB
}

func (databaseConnection *DatabaseConnection) MigrateDb() {
	databaseConnection.DB.AutoMigrate(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &model.Order{}, &model.Payment{}, &model.Product{}, &model.Stock{}, &model.Unit{}, &model.User{})
}

func (databaseConnection *DatabaseConnection) DropDb() {
	databaseConnection.DB.DropTable(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &model.Order{}, &model.Payment{}, &model.Product{}, &model.Stock{}, &model.Unit{}, &model.User{})
}
