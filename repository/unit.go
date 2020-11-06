package repository

import (
	"fmt"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type UnitRepository struct {
	db dbIface
}

func NewUnitRepo(db dbIface) *UnitRepository {
	return &UnitRepository{
		db: db,
	}
}

func (repo *UnitRepository) FindAll() ([]model.Unit, error) {
	var categories []model.Unit
	db := repo.db.GetDb()
	return categories, db.Find(&categories).Error
}

func (repo *UnitRepository) FindById(id int64) (model.Unit, error) {
	var unit model.Unit
	db := repo.db.GetDb()
	return unit, db.Where("id = ?", id).First(&unit).Error
}

func (repo *UnitRepository) FindByName(name string) ([]model.Unit, error) {
	var units []model.Unit
	db := repo.db.GetDb()
	return units, db.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&units).Error
}

func (repo *UnitRepository) New(unit model.Unit) (model.Unit, error) {
	var err error
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(unit)
	if isNotExist {
		err = db.Create(&unit).Error
	}
	return unit, err
}

func (repo *UnitRepository) Update(unit model.Unit) (model.Unit, error) {
	var oldUnit model.Unit
	db := repo.db.GetDb()
	err := db.Where("id = ?", unit.ID).First(&oldUnit).Error
	if err != nil {
		return unit, err
	}
	oldUnit = unit
	return unit, db.Save(&oldUnit).Error
}

func (repo *UnitRepository) Delete(id int64) (model.Unit, error) {
	var unit model.Unit
	db := repo.db.GetDb()
	err := db.Where("id = ?", id).First(&unit).Error
	if err != nil {
		return unit, err
	}
	err = db.Delete(&unit).Error
	return unit, err
}
