package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
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

// GetProducts godoc
// @Summary Get Products based on query
// @Description Get Products based on query
// @Tags product
// @Produce json
// @Param searchBy query string false "name, unit, category, or code"
// @Param query query string false "If this empty, it will fetch all products"
// @Success 200 {object} response.ResponseMapper{data=[]model.Product} "Return array of product"
// @Failure 404 {object} response.ResponseMapper{data=string} "Return error with message"
// @Failure 500 {object} response.ResponseMapper{data=string} "Return error with message"
// @Router /products [get]
func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var products []model.Product
	var err error
	searchBy := r.URL.Query().Get("searchBy")
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Product: GetProductsHandler (/products)")
		products, err = c.product.GetAllProduct()
		if err != nil {
			renderJSONError(w, http.StatusInternalServerError, err, "Cannot get all products")
			return
		}
	} else {
		golog.Infof("GET - Product: GetProductsByNameHandler (/products?searchBy=%s&query=%s)", searchBy, query)
		if searchBy == "" || searchBy == "name" {
			products, err = c.product.GetProductsByName(query)
			if err != nil {
				renderJSONError(w, http.StatusNotFound, err, "Cannot get products by name")
				return
			}
		} else if searchBy == "unit" {
			products, err = c.product.GetProductsByUnitName(query)
			if err != nil {
				renderJSONError(w, http.StatusNotFound, err, "Cannot get products by unit")
				return
			}
		} else if searchBy == "category" {
			products, err = c.product.GetProductsByCategoryName(query)
			if err != nil {
				renderJSONError(w, http.StatusNotFound, err, "Cannot get products by category")
				return
			}
		} else if searchBy == "code" {
			products, err = c.product.GetProductsByCode(query)
			if err != nil {
				renderJSONError(w, http.StatusNotFound, err, "Cannot get products by code")
				return
			}
		} else {
			products = []model.Product{}
		}
	}

	renderJSONSuccess(w, http.StatusOK, products, "Success getting products")
}

func (c *ProductController) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("GET - Product: GetProductByIdHandler (/products/id/%d)", id)
	product, err := c.product.GetOneProduct(id)
	if err != nil {
		renderJSONError(w, http.StatusNotFound, err, fmt.Sprintf("Cannot find product with id: %d", id))
		return
	}

	renderJSONSuccess(w, http.StatusOK, product, fmt.Sprintf("Success getting product with id: %d", id))
}

func (c *ProductController) GetProductByCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	code := vars["code"]
	golog.Infof("GET - Product: GetProductByCodeHandler (/products/code/%s)", code)
	product, err := c.product.GetOneProductByCode(code)
	if err != nil {
		renderJSONError(w, http.StatusNotFound, err, fmt.Sprintf("Cannot find product with code: %s", code))
		return
	}

	renderJSONSuccess(w, http.StatusOK, product, fmt.Sprintf("Success getting product with code: %s", code))
}

func (c *ProductController) NewProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Product: NewProductHandler (/products)")

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

	product, err := c.product.NewProduct(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusCreated, product, "Product created")
}

func (c *ProductController) ExportAllProductsToExcelHandler(w http.ResponseWriter, r *http.Request) {
	excel, err := c.product.ExportAllProductsToExcel()
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot export all products to excel")
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
		renderJSONError(w, http.StatusInternalServerError, err, "Cannot export all products to excel")
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	sheetName := vars["sheetName"]
	golog.Infof("POST - Product: UploadExcelProductHandler (/products/upload_excel/%s)", sheetName)

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

	err = r.ParseMultipartForm(20 << 20)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}
	defer file.Close()

	rowsLength, err := c.product.NewProductUsingExcel(sheetName, file)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	message := fmt.Sprintf("In progress for improting %d data from excel", rowsLength)
	renderJSONSuccess(w, http.StatusCreated, message, message)
}

func (c *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Product: UpdateProductHandler (/products/%d)", id)

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

	product, err := c.product.UpdateProduct(string(body))
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	renderJSONSuccess(w, http.StatusOK, product, "Product updated")
}

func (c *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	golog.Infof("DELETE - Product: DeleteProductHandler (/products/%d)", id)

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

	product, err := c.product.DeleteProduct(id)
	if err != nil {
		renderJSONError(w, http.StatusBadRequest, err, err.Error())
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderJSONSuccess(w, http.StatusOK, product, "Product deleted")
}
