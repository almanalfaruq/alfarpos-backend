package service

import (
	"encoding/json"
	"errors"
	"strings"

	"../model"
	"../repository"
)

type UnitService struct {
	repository.IUnitRepository
}

type IUnitService interface {
	GetAllUnit() ([]model.Unit, error)
	GetOneUnit(id int) (model.Unit, error)
	GetUnitsByName(name string) ([]model.Unit, error)
	NewUnit(unitData string) (model.Unit, error)
	UpdateUnit(unitData string) (model.Unit, error)
	DeleteUnit(id int) (model.Unit, error)
}

func (service *UnitService) GetAllUnit() ([]model.Unit, error) {
	return service.FindAll(), nil
}

func (service *UnitService) GetOneUnit(id int) (model.Unit, error) {
	unit := service.FindById(id)
	if unit.ID == 0 {
		return unit, errors.New("Unit not found")
	}
	return unit, nil
}

func (service *UnitService) GetUnitsByName(name string) ([]model.Unit, error) {
	name = strings.ToLower(name)
	units := service.FindByName(name)
	if len(units) == 0 {
		return units, errors.New("Units not found")
	}
	return units, nil
}

func (service *UnitService) NewUnit(unitData string) (model.Unit, error) {
	var unit model.Unit
	unitDataByte := []byte(unitData)
	err := json.Unmarshal(unitDataByte, &unit)
	if err != nil {
		return unit, err
	}
	return service.New(unit), nil
}

func (service *UnitService) UpdateUnit(unitData string) (model.Unit, error) {
	var unit model.Unit
	unitDataByte := []byte(unitData)
	err := json.Unmarshal(unitDataByte, &unit)
	if err != nil {
		return unit, err
	}
	unit = service.Update(unit)
	return unit, nil
}

func (service *UnitService) DeleteUnit(id int) (model.Unit, error) {
	return service.Delete(id)
}
