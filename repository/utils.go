package repository

import (
	"github.com/almanalfaruq/alfarpos-backend/util"
	"gorm.io/gorm"
)

type dbIface interface {
	Open(config util.Config) *gorm.DB
	GetDb() *gorm.DB
	MigrateDb()
	DropDb()
	Close()
}
