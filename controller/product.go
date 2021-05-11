package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/almanalfaruq/alfarpos-backend/model"
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/logger"
	"github.com/almanalfaruq/alfarpos-backend/util/response"
)

type ProductController struct {
	product productServiceIface
	conf    util.Config
}

func NewProductController(conf util.Config, productService productServiceIface) *ProductController {
	return &ProductController{
		conf:    conf,
		product: productService,
	}
}

// GetProductsByQuery godoc
// @Summary Get Products based on query
// @Description Get Products based on query
// @Tags product
// @Produce json
// @Param searchBy query string false "unit or category"
// @Param query query string false "If this empty, it will fetch all products"
// @Success 200 {object} response.ResponseMapper{data=[]model.Product} "Return array of product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products [get]
func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var (
		products []model.Product
		hasNext  bool
		err      error
	)
	searchBy := r.URL.Query().Get("searchBy")
	query := r.URL.Query().Get("query")
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

	if query == "" {
		logger.Log.Info("GET - Product: GetProductsHandler (/products)")
		products, hasNext, err = c.product.GetAllProduct(int(limit), int(page))
		if err != nil {
			response.RenderJSONError(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		logger.Log.Infof("GET - Product: GetProductsByNameHandler (/products?searchBy=%s&query=%s&limit=%d&page=%d)", searchBy, query, limit, page)
		if searchBy == "" {
			products, hasNext, err = c.product.GetProductsBySearchQuery(query, int(limit), int(page))
			if err != nil {
				response.RenderJSONError(w, http.StatusNotFound, err)
				return
			}
		} else if searchBy == "unit" {
			products, err = c.product.GetProductsByUnitName(query)
			if err != nil {
				response.RenderJSONError(w, http.StatusNotFound, err)
				return
			}
		} else if searchBy == "category" {
			products, err = c.product.GetProductsByCategoryName(query)
			if err != nil {
				response.RenderJSONError(w, http.StatusNotFound, err)
				return
			}
		} else {
			products = []model.Product{}
		}
	}

	resp := model.ProductResponse{
		Products: products,
		HasNext:  hasNext,
	}

	response.RenderJSONSuccess(w, http.StatusOK, resp, "Success getting products")
}

// GetProductsByID godoc
// @Summary Get Product based on id
// @Description Get Product based on id
// @Tags product
// @Produce json
// @Param id path integer false "id of the product"
// @Success 200 {object} response.ResponseMapper{data=model.Product} "Return a product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products/id/{id} [get]
func (c *ProductController) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("GET - Product: GetProductByIdHandler (/products/id/%d)", id)
	product, err := c.product.GetOneProduct(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, product, fmt.Sprintf("Success getting product with id: %d", id))
}

// GetProductsByIDs godoc
// @Summary Get Multiple Product based on ids
// @Description Get Multiple Product based on id
// @Tags product
// @Produce json
// @Param id path integer false "id of the product"
// @Success 200 {object} response.ResponseMapper{data=[]model.Product} "Return a product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products/id/{id} [get]
func (c *ProductController) GetProductsByIDsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	ids := vars["ids"]
	splitIDs := strings.Split(ids, ",")
	logger.Log.Infof("GET - Product: GetProductsByIDsHandler (/products/ids/%s)", ids)
	var intIDs []int64
	for _, id := range splitIDs {
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil || intID == 0 {
			continue
		}
		intIDs = append(intIDs, intID)
	}
	products, err := c.product.GetProductsByIDs(intIDs)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, products, fmt.Sprintf("Success getting products by ids: %s", ids))
}

// GetProductsByCode godoc
// @Summary Get Product based on code
// @Description Get Product based on code
// @Tags product
// @Produce json
// @Param coded path string false "code of the product"
// @Success 200 {object} response.ResponseMapper{data=model.Product} "Return a product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products/code/{code} [get]
func (c *ProductController) GetProductByCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	code := vars["code"]
	logger.Log.Infof("GET - Product: GetProductByCodeHandler (/products/code/%s)", code)
	product, err := c.product.GetOneProductByCode(code)
	if err != nil {
		response.RenderJSONError(w, http.StatusNotFound, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, product, fmt.Sprintf("Success getting product with code: %s", code))
}

func (c *ProductController) NewProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logger.Log.Info("POST - Product: NewProductHandler (/products)")

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

	product, err := c.product.NewProduct(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusCreated, product, "Product created")
}

func (c *ProductController) ExportAllProductsToExcelHandler(w http.ResponseWriter, r *http.Request) {
	excel, err := c.product.ExportAllProductsToExcel()
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	now := time.Now()
	excelName := fmt.Sprintf("Exported-Product-%02d-%02d-%d.xlsx", now.Day(), now.Month(), now.Year())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+excelName)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = excel.Write(w)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
	}
}

// UploadExcelProduct godoc
// @Summary Upload products from excel file
// @Description Upload products from excel file (*.xlsx). First row should be the header with these column in order:
// @Description Code (Barcode), Product Name, Sell Price, Quantity, Category Name, Buy Price, and Unit Name (pcs, bottle, etc.)
// @Tags product
// @Produce json
// @Param sheetName path string false "Name of the sheet"
// @Success 200 {object} response.ResponseMapper{data=[]model.Product} "Return array of product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products/upload_excel [post]
func (c *ProductController) UploadExcelProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	sheetName := vars["sheetName"]
	logger.Log.Infof("POST - Product: UploadExcelProductHandler (/products/upload_excel/%s)", sheetName)

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

	err = r.ParseMultipartForm(20 << 20)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	rowsLength, err := c.product.NewProductUsingExcel(sheetName, file)
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	message := fmt.Sprintf("In progress for improting %d data from excel", rowsLength)
	response.RenderJSONSuccess(w, http.StatusCreated, message, message)
}

func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	logger.Log.Infof("PUT - Product: UpdateProductHandler (/products/%d)", id)

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

	product, err := c.product.UpdateProduct(string(body))
	if err != nil {
		response.RenderJSONError(w, http.StatusInternalServerError, err)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, product, "Product updated")
}

func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	logger.Log.Infof("DELETE - Product: DeleteProductHandler (/products/%d)", id)

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

	product, err := c.product.DeleteProduct(id)
	if err != nil {
		response.RenderJSONError(w, http.StatusBadRequest, err)
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.RenderJSONSuccess(w, http.StatusOK, product, "Product deleted")
}
