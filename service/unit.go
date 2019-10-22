package service

import (
	"encoding/json"

	"../model"
	"../repository"
)

type UnitService struct {
	repository.IUnitRepository
}

type IUnitService interface {
	GetAllUnit() []model.Unit
	GetOneUnit(id int) (model.Unit, error)
	NewUnit(unitData string) (model.Unit, error)
	UpdateUnit(unitData string) (model.Unit, error)
	DeleteUnit(id int) int
}

func (service *UnitService) GetAllUnit() []model.Unit {
	return service.FindAll()
}

func (service *UnitService) GetOneUnit(id int) (model.Unit, error) {
	return service.FindById(id), nil
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
	return service.Delete(id), nil
}
