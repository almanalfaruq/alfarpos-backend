package service_test

import (
	"testing"

	"github.com/almanalfaruq/alfarpos-backend/model"
	. "github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/test/mocks"
	"github.com/almanalfaruq/alfarpos-backend/test/resources"
	"github.com/stretchr/testify/assert"
)

func TestGetUnitById(t *testing.T) {
	t.Run("Get Unit By ID - Pass", func(t *testing.T) {
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("FindById", 1).Return(resources.Unit1)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := resources.Unit1

		actualResult, err := unitService.GetOneUnit(1)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Get Unit By ID - Error", func(t *testing.T) {
		var unit model.Unit
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("FindById", 5).Return(unit)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := unit

		actualResult, err := unitService.GetOneUnit(5)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestGetAllUnit(t *testing.T) {
	unitRepository := new(mocks.UnitRepository)

	unitRepository.On("FindAll").Return(resources.Units)

	unitService := UnitService{
		IUnitRepository: unitRepository,
	}

	expectedResult := resources.Units

	actualResult, err := unitService.GetAllUnit()

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult[0].ID, actualResult[0].ID)
	assert.Equal(t, expectedResult[0].Name, actualResult[0].Name)
	assert.Equal(t, expectedResult[1].ID, actualResult[1].ID)
	assert.Equal(t, expectedResult[1].Name, actualResult[1].Name)
	assert.Equal(t, expectedResult[2].ID, actualResult[2].ID)
	assert.Equal(t, expectedResult[2].Name, actualResult[2].Name)
	assert.Equal(t, expectedResult[3].ID, actualResult[3].ID)
	assert.Equal(t, expectedResult[3].Name, actualResult[3].Name)
}

func TestNewUnit(t *testing.T) {
	t.Run("New Unit - Success", func(t *testing.T) {
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("New", resources.Unit2).Return(resources.Unit2)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := resources.Unit2

		jsonUnit := `{"id": 2, "name": "Unit2"}`

		actualResult, err := unitService.NewUnit(jsonUnit)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("New Unit - Error", func(t *testing.T) {
		var unit model.Unit
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("New", unit).Return(unit)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := unit

		jsonUnit := `{name: "Unit2"}`

		actualResult, err := unitService.NewUnit(jsonUnit)

		assert.NotNil(t, err)
		assert.NotNil(t, actualResult)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestUpdateUnit(t *testing.T) {
	t.Run("Update Unit - Success", func(t *testing.T) {
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("Update", resources.Unit2Updated).Return(resources.Unit2Updated)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := resources.Unit2Updated

		jsonUnit := `{"id": 2, "name": "Unit2Updated"}`

		actualResult, err := unitService.UpdateUnit(jsonUnit)

		assert.Nil(t, err)
		assert.NotNil(t, actualResult)
		assert.NotEmpty(t, actualResult)
		assert.Equal(t, expectedResult.ID, actualResult.ID)
		assert.Equal(t, expectedResult.Name, actualResult.Name)
	})

	t.Run("Update Unit - Error", func(t *testing.T) {
		var unit model.Unit
		unitRepository := new(mocks.UnitRepository)

		unitRepository.On("Update", unit).Return(unit)

		unitService := UnitService{
			IUnitRepository: unitRepository,
		}

		expectedResult := unit

		jsonUnit := `{name: "Unit2Updated"}`

		actualResult, err := unitService.UpdateUnit(jsonUnit)

		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, actualResult)
	})
}

func TestDeleteUnit(t *testing.T) {
	unitRepository := new(mocks.UnitRepository)

	unitRepository.On("Delete", 4).Return(resources.Unit4)

	unitService := UnitService{
		IUnitRepository: unitRepository,
	}

	expectedResult := resources.Unit4

	id := 4

	actualResult, err := unitService.DeleteUnit(id)

	assert.Nil(t, err)
	assert.NotNil(t, actualResult)
	assert.NotEmpty(t, actualResult)
	assert.Equal(t, expectedResult.ID, actualResult.ID)
	assert.Equal(t, expectedResult.Name, actualResult.Name)
}
