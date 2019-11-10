package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	"github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
)

type CategoryController struct {
	service.ICategoryService
	util.Config
}

func (controller *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var categories []model.Category
	var err error
	var responseMapper response.ResponseMapper
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Category: GetAllCategoryHandler (/categories)")
		categories, err = controller.GetAllCategory()
		if err != nil {
			golog.Error(err)
			responseMapper = response.ResponseMapper{
				Code:    http.StatusInternalServerError,
				Data:    err.Error(),
				Message: "Cannot get all categories",
			}
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(responseMapper)
			if err != nil {
				golog.Error("Cannot encode json")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	} else {
		golog.Infof("GET - Product: GetCategoriesByNameHandler (/categories?query=%s)", query)
		categories, err = controller.GetCategoriesByName(query)
		if err != nil {
			golog.Error(err)
			responseMapper = response.ResponseMapper{
				Code:    http.StatusNotFound,
				Data:    err.Error(),
				Message: "Cannot get categories by name",
			}
			w.WriteHeader(http.StatusNotFound)
			err = json.NewEncoder(w).Encode(responseMapper)
			if err != nil {
				golog.Error("Cannot encode json")
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	responseMapper = response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    categories,
		Message: "Success getting categories",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CategoryController) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var responseMapper response.ResponseMapper
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("GET - Product: GetCategoryByIdHandler (/categories/id/%v)", id)
	category, err := controller.GetOneCategory(int(id))
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusNotFound,
			Data:    err.Error(),
			Message: fmt.Sprintf("Cannot find category with id: %v", id),
		}
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error("Cannot encode json")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	responseMapper = response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    category,
		Message: fmt.Sprintf("Success getting category with id: %v", id),
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CategoryController) NewCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Category: NewCategoryHandler (/categories)")

	authHeader := r.Header.Get("Authorization")
	user, err := ParseJwtToUser(authHeader, controller.SecretKey)

	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusBadRequest,
			Data:    err.Error(),
			Message: "Cannot parse token",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	if user.RoleID == model.Cashier {
		golog.Error("User must be Admin or Manager")
		w.WriteHeader(http.StatusForbidden)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusForbidden,
			Data:    "User must be Admin or Manager",
			Message: "User must be Admin or Manager",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot read request body",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	category, err := controller.NewCategory(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Category: UpdateCategoryHandler (/categories/%v)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := ParseJwtToUser(authHeader, controller.SecretKey)

	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusBadRequest,
			Data:    err.Error(),
			Message: "Cannot parse token",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	if user.RoleID == model.Cashier {
		golog.Error("User must be Admin or Manager")
		w.WriteHeader(http.StatusForbidden)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusForbidden,
			Data:    "User must be Admin or Manager",
			Message: "User must be Admin or Manager",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusInternalServerError,
			Data:    err.Error(),
			Message: "Cannot read request body",
		}
		err = json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	category, err := controller.UpdateCategory(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("DELETE - Category: DeleteCategoryHandler (/categories/%v)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := ParseJwtToUser(authHeader, controller.SecretKey)

	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusBadRequest,
			Data:    err.Error(),
			Message: "Cannot parse token",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	if user.RoleID == model.Cashier {
		golog.Error("User must be Admin or Manager")
		w.WriteHeader(http.StatusForbidden)
		responseMapper := response.ResponseMapper{
			Code:    http.StatusForbidden,
			Data:    "User must be Admin or Manager",
			Message: "User must be Admin or Manager",
		}
		err := json.NewEncoder(w).Encode(responseMapper)
		if err != nil {
			golog.Error(err)
			http.Error(w, err.Error(), 500)
		}
		return
	}

	category, err := controller.DeleteCategory(int(id))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
