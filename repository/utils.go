package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/jinzhu/gorm"
)

type dbIface interface {
	Open(config util.Config) *gorm.DB
	GetDb() *gorm.DB
	MigrateDb()
	DropDb()
	Close()
}
