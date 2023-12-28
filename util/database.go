package util

import (
	"fmt"

	"net/url"

	"github.com/almanalfaruq/alfarpos-backend/model"
	orderentity "github.com/almanalfaruq/alfarpos-backend/model/order"
	productentity "github.com/almanalfaruq/alfarpos-backend/model/product"
	profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"
	statsentity "github.com/almanalfaruq/alfarpos-backend/model/stats"
	stockentity "github.com/almanalfaruq/alfarpos-backend/model/stock"
	transactionentity "github.com/almanalfaruq/alfarpos-backend/model/transaction"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/kataras/golog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	DB *gorm.DB
}

func (dbConn *DBConn) Open(config Config) *gorm.DB {
	var dbName string
	if config.Env == "test" {
		dbName = config.Database.DBTestName
		golog.Infof("Connecting to database: %v", config.Database.DBTestName)
	} else {
		dbName = config.Database.DBName
		golog.Infof("Connecting to database: %v", config.Database.DBName)
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.Username, url.QueryEscape(config.Database.Password),
		config.Database.Host, config.Database.Port, dbName)
	var err error
	dbConn.DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		golog.Fatalf("Cannot connect to the database using gorm. err: %v", err)
	}
	return dbConn.DB
}

func (dbConn *DBConn) GetDb() *gorm.DB {
	return dbConn.DB
}

func (dbConn *DBConn) MigrateDb() {
	err := dbConn.DB.AutoMigrate(&model.Category{}, &model.Customer{}, &orderentity.Order{}, &orderentity.OrderDetail{}, &model.Payment{},
		&productentity.Product{}, &productentity.ProductPrice{}, &stockentity.Stock{}, &model.Unit{}, &userentity.User{}, &transactionentity.Money{},
		&profileentity.Profile{}, &statsentity.ShopStats{})
	if err != nil {
		golog.Fatalf("Cannot connect to the database using gorm. err: %v", err)
	}
}

func (dbConn *DBConn) DropDb() {
	err := dbConn.DB.Migrator().DropTable(&model.Category{}, &model.Customer{}, &orderentity.Order{}, &orderentity.OrderDetail{}, &model.Payment{},
		&productentity.Product{}, &productentity.ProductPrice{}, &stockentity.Stock{}, &model.Unit{}, &userentity.User{}, &transactionentity.Money{},
		&profileentity.Profile{}, &statsentity.ShopStats{})
	if err != nil {
		golog.Fatalf("Cannot connect to the database using gorm. err: %v", err)
	}
}

func (dbConn *DBConn) Close() {
	db, _ := dbConn.DB.DB()
	db.Close()
}

type DBIface interface {
	Open(config Config) *gorm.DB
	GetDb() *gorm.DB
	MigrateDb()
	DropDb()
	Close()
}
