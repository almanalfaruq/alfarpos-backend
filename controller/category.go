package controller

import (
	"encoding/json"
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

// GetCategories godoc
// @Summary Get all categories or search by name
// @Description Get all categories or search by name
// @Tags category
// @Produce json
// @Param query query string false "Category name query"
// @Success 200 {object} response.ResponseStruct{data=[]model.Category} "Return array of categories"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /categories [get]
func (c *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var categories []model.Category
	var err error
	query := r.URL.Query().Get("query")

	if query == "" {
		logger.Log.Info("GET - Category: GetAllCategoryHandler (/categories)")
		categories, err = c.category.GetAllCategory()
		if err != nil {
			logger.Log.Error(err)
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		logger.Log.Infof("GET - Product: GetCategoriesByNameHandler (/categories?query=%s)", query)
		categories, err = c.category.GetCategoriesByName(query)
		if err != nil {
			response.RenderJSONError(w, http.StatusNotFound, err)
			return
		}
	}

	message := "Success getting categories"
	response.RenderJSONSuccess(w, http.StatusOK, categories, message)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get category by ID
// @Tags category
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ResponseStruct{data=model.Category} "Return category"
// @Failure 404 {object} response.ResponseStruct{data=string} "Not Found"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Router /categories/id/{id} [get]
func (c *CategoryController) GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("GET - Product: GetCategoryByIdHandler (/categories/id/%v)", id)
	category, err := c.category.GetOneCategory(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	message := fmt.Sprintf("Success getting category with id: %v", id)
	response.RenderJSONSuccess(w, http.StatusOK, category, message)
}

// NewCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags category
// @Accept json
// @Produce json
// @Param category body model.Category true "Category data"
// @Success 201 {object} response.ResponseStruct{data=model.Category} "Category created"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /categories [post]
func (c *CategoryController) NewCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - Category: NewCategoryHandler (/categories)")

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

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body model.Category true "Category data"
// @Success 200 {object} response.ResponseStruct{data=model.Category} "Category updated"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /categories/{id} [put]
func (c *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("PUT - Category: UpdateCategoryHandler (/categories/%v)", id)

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

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags category
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.ResponseStruct{data=model.Category} "Category deleted"
// @Failure 403 {object} response.ResponseStruct{data=string} "Forbidden"
// @Failure 500 {object} response.ResponseStruct{data=string} "Internal Server Error"
// @Security ApiKeyAuth
// @Router /categories/{id} [delete]
func (c *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("DELETE - Category: DeleteCategoryHandler (/categories/%v)", id)

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

	category, err := c.category.DeleteCategory(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, category, "Category updated")
}
