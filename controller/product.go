package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/almanalfaruq/alfarpos-backend/model"
	"github.com/almanalfaruq/alfarpos-backend/model/response"
	"github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
)

type ProductController struct {
	service.IProductService
	util.Config
}

func (controller *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var products []model.Product
	var err error
	var responseMapper response.ResponseMapper
	searchBy := r.URL.Query().Get("searchBy")
	query := r.URL.Query().Get("query")

	if query == "" {
		golog.Info("GET - Product: GetProductsHandler (/products)")
		products, err = controller.GetAllProduct()
		if err != nil {
			golog.Error(err)
			responseMapper = response.ResponseMapper{
				Code:    http.StatusInternalServerError,
				Data:    err.Error(),
				Message: "Cannot get all products",
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
		golog.Infof("GET - Product: GetProductsByNameHandler (/products?searchBy=%s&query=%s)", searchBy, query)
		if searchBy == "" || searchBy == "name" {
			products, err = controller.GetProductsByName(query)
			if err != nil {
				golog.Error(err)
				responseMapper = response.ResponseMapper{
					Code:    http.StatusNotFound,
					Data:    err.Error(),
					Message: "Cannot get products by name",
				}
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(responseMapper)
				if err != nil {
					golog.Error("Cannot encode json")
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		} else if searchBy == "unit" {
			products, err = controller.GetProductsByUnitName(query)
			if err != nil {
				golog.Error(err)
				responseMapper = response.ResponseMapper{
					Code:    http.StatusNotFound,
					Data:    err.Error(),
					Message: "Cannot get products by unit",
				}
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(responseMapper)
				if err != nil {
					golog.Error("Cannot encode json")
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		} else if searchBy == "category" {
			products, err = controller.GetProductsByCategoryName(query)
			if err != nil {
				golog.Error(err)
				responseMapper = response.ResponseMapper{
					Code:    http.StatusNotFound,
					Data:    err.Error(),
					Message: "Cannot get products by category",
				}
				w.WriteHeader(http.StatusNotFound)
				err = json.NewEncoder(w).Encode(responseMapper)
				if err != nil {
					golog.Error("Cannot encode json")
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		} else if searchBy == "code" {
			products, err = controller.GetProductsByCode(query)
			if err != nil {
				golog.Error(err)
				responseMapper = response.ResponseMapper{
					Code:    http.StatusNotFound,
					Data:    err.Error(),
					Message: "Cannot get products by code",
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
			products = []model.Product{}
		}
	}

	responseMapper = response.ResponseMapper{
		Code:    http.StatusOK,
		Data:    products,
		Message: "Success getting products",
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var responseMapper response.ResponseMapper
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("GET - Product: GetProductByIdHandler (/products/id/%v)", id)
	product, err := controller.GetOneProduct(int(id))
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusNotFound,
			Data:    err.Error(),
			Message: fmt.Sprintf("Cannot find product with id: %v", id),
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
		Data:    product,
		Message: fmt.Sprintf("Success getting product with id: %v", id),
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) GetProductByCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var responseMapper response.ResponseMapper
	vars := mux.Vars(r)
	code := vars["code"]
	golog.Infof("GET - Product: GetProductByCodeHandler (/products/code/%v)", code)
	product, err := controller.GetOneProductByCode(code)
	if err != nil {
		golog.Error(err)
		responseMapper = response.ResponseMapper{
			Code:    http.StatusNotFound,
			Data:    err.Error(),
			Message: fmt.Sprintf("Cannot find product with code: %v", code),
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
		Data:    product,
		Message: fmt.Sprintf("Success getting product with code: %v", code),
	}
	err = json.NewEncoder(w).Encode(responseMapper)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) NewProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	golog.Info("POST - Product: NewProductHandler (/products)")

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

	product, err := controller.NewProduct(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) UploadExcelProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	sheetName := vars["sheetName"]
	golog.Infof("POST - Product: UploadExcelProductHandler (/products/upload_excel/%v)", sheetName)

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

	err = r.ParseMultipartForm(20 << 20)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = controller.NewProductUsingExcel(sheetName, file)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products, err := controller.GetAllProduct()
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("PUT - Product: UpdateProductHandler (/products/%v)", id)

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

	product, err := controller.UpdateProduct(string(body))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *ProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	golog.Infof("DELETE - Product: DeleteProductHandler (/products/%v)", id)

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

	product, err := controller.DeleteProduct(int(id))
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
