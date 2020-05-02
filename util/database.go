package util

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

type DBConn struct {
	DB *gorm.DB
}

func (dbConn *DBConn) Open(config Config) *gorm.DB {
	var url string
	if config.Env == "test" {
		golog.Infof("Connecting to database: %v", config.Database.DBTestName)
		url = fmt.Sprintf("host=%v port=%v user=%v password=\"%v\" dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBTestName)
	} else {
		golog.Infof("Connecting to database: %v", config.Database.DBName)
		url = fmt.Sprintf("host=%v port=%v user=%v password=\"%v\" dbname=%v sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName)
	}
	golog.Infof("URL Database: %s", url)
	var err error
	dbConn.DB, err = gorm.Open("postgres", url)
	if err != nil {
		panic("Cannot connect to the database")
	}
	return dbConn.DB
}

func (dbConn *DBConn) GetDb() *gorm.DB {
	return dbConn.DB
}

func (dbConn *DBConn) MigrateDb() {
	dbConn.DB.AutoMigrate(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &model.Order{}, &model.Payment{}, &model.Product{}, &model.Stock{}, &model.Unit{}, &model.User{})
}

func (dbConn *DBConn) DropDb() {
	dbConn.DB.DropTable(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &model.Order{}, &model.Payment{}, &model.Product{}, &model.Stock{}, &model.Unit{}, &model.User{})
}

func (dbConn *DBConn) Close() {
	dbConn.DB.Close()
}
