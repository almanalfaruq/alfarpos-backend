package repository

import (
	"../model"
	"../util"
)

type UnitRepository struct {
	util.IDatabaseConnection
}

type IUnitRepository interface {
	FindAll() []model.Unit
	FindById(id int) model.Unit
	FindByName(name string) model.Unit
	New(unit model.Unit) model.Unit
	Update(unit model.Unit) model.Unit
	Delete(id int) model.Unit
}

func (repo *UnitRepository) FindAll() []model.Unit {
	var categories []model.Unit
	db := repo.GetDb()
	db.Find(&categories)
	return categories
}

func (repo *UnitRepository) FindById(id int) model.Unit {
	var unit model.Unit
	db := repo.GetDb()
	db.Where("id = ?", id).First(&unit)
	return unit
}

func (repo *UnitRepository) FindByName(name string) model.Unit {
	var unit model.Unit
	db := repo.GetDb()
	db.Where("name = ?", name).First(&unit)
	return unit
}

func (repo *UnitRepository) New(unit model.Unit) model.Unit {
	db := repo.GetDb()
	isNotExist := db.NewRecord(unit)
	if isNotExist {
		db.Create(&unit)
	}
	return unit
}

func (repo *UnitRepository) Update(unit model.Unit) model.Unit {
	var oldUnit model.Unit
	db := repo.GetDb()
	db.Where("id = ?", unit.ID).First(&oldUnit)
	oldUnit = unit
	db.Save(&oldUnit)
	return unit
}

func (repo *UnitRepository) Delete(id int) model.Unit {
	var unit model.Unit
	db := repo.GetDb()
	db.Where("id = ?", id).First(&unit)
	db.Delete(&unit)
	return unit
}
