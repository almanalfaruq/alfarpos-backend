package util

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/kataras/golog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	DB *gorm.DB
}

func (dbConn *DBConn) Open(config Config) *gorm.DB {
	var (
		dbName   string
		password string
		url      string
	)
	if config.Env == "test" {
		dbName = config.Database.DBTestName
		golog.Infof("Connecting to database: %v", config.Database.DBTestName)
	} else {
		dbName = config.Database.DBName
		golog.Infof("Connecting to database: %v", config.Database.DBName)
	}
	if config.Database.Password == "" {
		password = `""`
	} else {
		password = config.Database.Password
	}
	url = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port,
		config.Database.Username, password, dbName)
	golog.Infof("URL Database: %s", url)
	var err error
	dbConn.DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to the database")
	}
	return dbConn.DB
}

func (dbConn *DBConn) GetDb() *gorm.DB {
	return dbConn.DB
}

func (dbConn *DBConn) MigrateDb() {
	err := dbConn.DB.AutoMigrate(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &orderentity.Order{}, &model.Payment{}, &model.Product{}, &model.ProductPrice{}, &model.Stock{}, &model.Unit{}, &userentity.User{})
	if err != nil {
		panic(err)
	}
}

func (dbConn *DBConn) DropDb() {
	dbConn.DB.Migrator().DropTable(&model.Category{}, &model.Customer{}, &model.OrderDetail{}, &orderentity.Order{}, &model.Payment{}, &model.Product{}, &model.ProductPrice{}, &model.Stock{}, &model.Unit{}, &userentity.User{})
}

func (dbConn *DBConn) Close() {
	db, _ := dbConn.DB.DB()
	db.Close()
}
