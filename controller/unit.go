package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
)

type UnitController struct {
	unit unitServiceIface
	conf util.Config
}

func NewUnitController(conf util.Config, unitService unitServiceIface) *UnitController {
	return &UnitController{
		unit: unitService,
		conf: conf,
	}
}

// GetUnits godoc
// @Summary Get all units or search by name
// @Description Get all units or search by name
// @Tags unit
// @Produce json
// @Param query query string false "Unit name query"
// @Success 200 {object} response.ResponseStruct{data=[]model.Unit} "Return array of units"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /units [get]
func (c *UnitController) GetUnitsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var units []model.Unit
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		logger.Log.Info("GET - Unit: GetAllUnitHandler (/units)")
		units, err = c.unit.GetAllUnit()
		if err != nil {
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		logger.Log.Infof("GET - Product: GetUnitsByNameHandler (/units?query=%s)", query)
		units, err = c.unit.GetUnitsByName(query)
		if err != nil {
			response.RenderJSONError(w, http.StatusNotFound, err)
			return
		}
	}

	response.RenderJSONSuccess(w, http.StatusOK, units, "Success getting all units")
}

// GetUnitByID godoc
// @Summary Get unit by ID
// @Description Get unit by ID
// @Tags unit
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} response.ResponseStruct{data=model.Unit} "Return unit"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /units/id/{id} [get]
func (c *UnitController) GetUnitByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("GET - Product: GetUnitByIdHandler (/units/id/%d)", id)
	unit, err := c.unit.GetOneUnit(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, unit, fmt.Sprintf("Success getting unit with id: %d", id))
}

// NewUnit godoc
// @Summary Create a new unit
// @Description Create a new unit
// @Tags unit
// @Accept json
// @Produce json
// @Param unit body model.Unit true "Unit data"
// @Success 201 {object} response.ResponseStruct{data=model.Unit} "Unit created"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /units [post]
func (c *UnitController) NewUnitHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - Unit: NewUnitHandler (/units)")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	unit, err := c.unit.NewUnit(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, unit, "Unit created")
}

// UpdateUnit godoc
// @Summary Update a unit
// @Description Update a unit
// @Tags unit
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Param unit body model.Unit true "Unit data"
// @Success 200 {object} response.ResponseStruct{data=model.Unit} "Unit updated"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /units/{id} [put]
func (c *UnitController) UpdateUnitHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("PUT - Unit: UpdateUnitHandler (/units/%d)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	unit, err := c.unit.UpdateUnit(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, unit, "Unit updated")
}

// DeleteUnit godoc
// @Summary Delete a unit
// @Description Delete a unit by ID
// @Tags unit
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} response.ResponseStruct{data=model.Unit} "Unit deleted"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /units/{id} [delete]
func (c *UnitController) DeleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("DELETE - Unit: DeleteUnitHandler (/units/%d)", id)

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, errors.New(message))
		return
	}

	unit, err := c.unit.DeleteUnit(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, unit, "Unit deleted")
}
