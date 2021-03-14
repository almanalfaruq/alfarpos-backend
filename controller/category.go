package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
)

type CategoryController struct {
	category categoryServiceIface
	conf     util.Config
}

func NewCategoryController(conf util.Config, categoryService categoryServiceIface) *CategoryController {
	return &CategoryController{
		category: categoryService,
		conf:     conf,
	}
}

func (c *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var categories []model.Category
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Category: GetAllCategoryHandler (/categories)")
		categories, err = c.category.GetAllCategory()
		if err != nil {
			golog.Error(err)
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		golog.Infof("GET - Product: GetCategoriesByNameHandler (/categories?query=%s)", query)
		categories, err = c.category.GetCategoriesByName(query)
		if err != nil {
			response.RenderJSONError(w, http.StatusNotFound, err)
			return
		}
	}

	message := "Success getting categories"
	response.RenderJSONSuccess(w, http.StatusOK, categories, message)
}

func (c *CategoryController) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("GET - Product: GetCategoryByIdHandler (/categories/id/%v)", id)
	category, err := c.category.GetOneCategory(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	message := fmt.Sprintf("Success getting category with id: %v", id)
	response.RenderJSONSuccess(w, http.StatusOK, category, message)
}

func (c *CategoryController) NewCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Category: NewCategoryHandler (/categories)")

	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var data model.Category
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}
	category, err := c.category.NewCategory(data.Name)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, category, "Category created")
}

func (c *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("PUT - Category: UpdateCategoryHandler (/categories/%v)", id)

<<<<<<< HEAD
	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
=======
	user, ok := r.Context().Value(model.CTX_USER).(model.User)
>>>>>>> heroku
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	var data model.Category
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	category, err := c.category.UpdateCategory(data)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, category, "Category updated")
}

func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Category: DeleteCategoryHandler (/categories/%v)", id)

<<<<<<< HEAD
	user, ok := r.Context().Value(userentity.CTX_USER).(userentity.User)
=======
	user, ok := r.Context().Value(model.CTX_USER).(model.User)
>>>>>>> heroku
	if !ok {
		err := errors.New("Cannot parse user context")
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	if ok := user.HasRole(userentity.RoleManager, userentity.RoleAdmin); !ok {
		message := "User must be Admin or Manager"
		response.RenderJSONError(w, http.StatusForbidden, fmt.Errorf(message))
		return
	}

	category, err := c.category.DeleteCategory(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, category, "Category updated")
}
