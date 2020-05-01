package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
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
			errMessage := "Cannot get all categories"
			renderJSONError(w, http.StatusInternalServerError, err, errMessage)
			return
		}
	} else {
		golog.Infof("GET - Product: GetCategoriesByNameHandler (/categories?query=%s)", query)
		categories, err = c.category.GetCategoriesByName(query)
		if err != nil {
			errMessage := "Cannot get categories by name"
			renderJSONError(w, http.StatusNotFound, err, errMessage)
			return
		}
	}

	message := "Success getting categories"
	renderJSONSuccess(w, http.StatusOK, categories, message)
}

func (c *CategoryController) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("GET - Product: GetCategoryByIdHandler (/categories/id/%v)", id)
	category, err := c.category.GetOneCategory(id)
	if err != nil {
		errMessage := fmt.Sprintf("Cannot find category with id: %v", id)
		renderJSONError(w, http.StatusNotFound, err, errMessage)
		return
	}

	message := fmt.Sprintf("Success getting category with id: %v", id)
	renderJSONSuccess(w, http.StatusOK, category, message)
}

func (c *CategoryController) NewCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Category: NewCategoryHandler (/categories)")

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
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

	var data model.Category
	err = json.Unmarshal(body, &data)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot unmarshal JSON")
		return
	}
	category, err := c.category.NewCategory(data.Name)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, category, "Category created")
}

func (c *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("PUT - Category: UpdateCategoryHandler (/categories/%v)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
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

	var data model.Category
	err = json.Unmarshal(body, &data)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot unmarshal JSON")
		return
	}
	category, err := c.category.UpdateCategory(data)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, category, "Category updated")
}

func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Category: DeleteCategoryHandler (/categories/%v)", id)

	authHeader := r.Header.Get("Authorization")
	user, err := parseJwtToUser(authHeader, c.conf.SecretKey)

	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, "Cannot parse token")
		return
	}

	if ok := user.HasRole(model.Manager, model.Admin); !ok {
		message := "User must be Admin or Manager"
		renderJSONError(w, http.StatusForbidden, fmt.Errorf(message), message)
		return
	}

	category, err := c.category.DeleteCategory(id)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, category, "Category updated")
}
