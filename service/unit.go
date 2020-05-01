package service

import (
	"encoding/json"
	"errors"
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
	return service.unit.FindAll(), nil
}

func (service *UnitService) GetOneUnit(id int) (model.Unit, error) {
	unit := service.unit.FindById(id)
	if unit.ID == 0 {
		return unit, errors.New("Unit not found")
	}
	return unit, nil
}

func (service *UnitService) GetUnitsByName(name string) ([]model.Unit, error) {
	name = strings.ToLower(name)
	units := service.unit.FindByName(name)
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
	return service.unit.New(unit), nil
}

func (service *UnitService) UpdateUnit(unitData string) (model.Unit, error) {
	var unit model.Unit
	unitDataByte := []byte(unitData)
	err := json.Unmarshal(unitDataByte, &unit)
	if err != nil {
		return unit, err
	}
	unit = service.unit.Update(unit)
	return unit, nil
}

func (service *UnitService) DeleteUnit(id int) (model.Unit, error) {
	return service.unit.Delete(id)
}
