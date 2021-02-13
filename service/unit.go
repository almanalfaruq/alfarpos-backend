package service

import (
	"encoding/json"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/model"
)

type UnitService struct {
	unit unitRepositoryIface
}

func NewUnitService(unitRepo unitRepositoryIface) *UnitService {
	return &UnitService{
		unit: unitRepo,
	}
}

func (service *UnitService) GetAllUnit() ([]model.Unit, error) {
	return service.unit.FindAll()
}

func (service *UnitService) GetOneUnit(id int64) (model.Unit, error) {
	return service.unit.FindById(id)
}

func (service *UnitService) GetUnitsByName(name string) ([]model.Unit, error) {
	name = strings.ToLower(name)
	return service.unit.FindByName(name)
}

func (service *UnitService) NewUnit(unitData string) (model.Unit, error) {
	var unit model.Unit
	unitDataByte := []byte(unitData)
	err := json.Unmarshal(unitDataByte, &unit)
	if err != nil {
		return model.Unit{}, err
	}
	return service.unit.New(unit)
}

func (service *UnitService) UpdateUnit(unitData string) (model.Unit, error) {
	var unit model.Unit
	unitDataByte := []byte(unitData)
	err := json.Unmarshal(unitDataByte, &unit)
	if err != nil {
		return unit, err
	}
	return service.unit.Update(unit)
}

func (service *UnitService) DeleteUnit(id int64) (model.Unit, error) {
	return service.unit.Delete(id)
}
