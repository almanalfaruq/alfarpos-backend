package util

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

func (databaseConnection *DatabaseConnection) Open(config Config) *gorm.DB {
	var url string
	if config.Env == "test" {
		golog.Infof("Connecting to database: %v", config.Database.DBTestName)
		url = fmt.Sprintf("host=%v port=%v user=%v password=\"%v\" dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBTestName)
	} else if config.Env == "dev" {
		golog.Infof("Connecting to database: %v", config.Database.DBName)
		url = fmt.Sprintf("host=%v port=%v user=%v password=\"%v\" dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName)
	}
	golog.Infof("URL Database: %s", url)
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

func (databaseConnection *DatabaseConnection) Close() {
	databaseConnection.DB.Close()
}
