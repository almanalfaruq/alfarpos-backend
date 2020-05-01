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

func (repo *UnitRepository) FindAll() []model.Unit {
	var categories []model.Unit
	db := repo.db.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *UnitRepository) FindById(id int) model.Unit {
	var unit model.Unit
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&unit)
	return unit
}

func (repo *UnitRepository) FindByName(name string) []model.Unit {
	var units []model.Unit
	db := repo.db.GetDb()
	db.Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&units)
	return units
}

func (repo *UnitRepository) New(unit model.Unit) model.Unit {
	db := repo.db.GetDb()
	isNotExist := db.NewRecord(unit)
	if isNotExist {
		db.Create(&unit)
	}
	return unit
}

func (repo *UnitRepository) Update(unit model.Unit) model.Unit {
	var oldUnit model.Unit
	db := repo.db.GetDb()
	db.Where("id = ?", unit.ID).First(&oldUnit)
	oldUnit = unit
	db.Save(&oldUnit)
	return unit
}

func (repo *UnitRepository) Delete(id int) (model.Unit, error) {
	var unit model.Unit
	db := repo.db.GetDb()
	db.Where("id = ?", id).First(&unit)
	err := db.Delete(&unit).Error
	return unit, err
}
