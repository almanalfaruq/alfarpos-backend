package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
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

func (c *UnitController) GetUnitsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var units []model.Unit
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Unit: GetAllUnitHandler (/units)")
		units, err = c.unit.GetAllUnit()
		if err != nil {
			renderJSONError(w, http.StatusInternalServerError, err, "Cannot get all units")
			return
		}
	} else {
		golog.Infof("GET - Product: GetUnitsByNameHandler (/units?query=%s)", query)
		units, err = c.unit.GetUnitsByName(query)
		if err != nil {
			renderJSONError(w, http.StatusNotFound, err, "Cannot get units by name")
			return
		}
	}

	renderJSONSuccess(w, http.StatusOK, units, "Success getting all units")
}

func (c *UnitController) GetUnitByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("GET - Product: GetUnitByIdHandler (/units/id/%d)", id)
	unit, err := c.unit.GetOneUnit(id)
	if err != nil {
		renderJSON(w, http.StatusNotFound, err, fmt.Sprintf("Cannot find unit with id: %d", id))
		return
	}

	renderJSONSuccess(w, http.StatusOK, unit, fmt.Sprintf("Success getting unit with id: %d", id))
}

func (c *UnitController) NewUnitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Unit: NewUnitHandler (/units)")

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.RoleManager, model.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	unit, err := c.unit.NewUnit(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, unit, "Unit created")
}

func (c *UnitController) UpdateUnitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("PUT - Unit: UpdateUnitHandler (/units/%d)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.RoleManager, model.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot read request body")
		return
	}

	unit, err := c.unit.UpdateUnit(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusOK, unit, "Unit updated")
}

func (c *UnitController) DeleteUnitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Unit: DeleteUnitHandler (/units/%d)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.RoleManager, model.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	unit, err := c.unit.DeleteUnit(id)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusOK, unit, "Unit deleted")
}
