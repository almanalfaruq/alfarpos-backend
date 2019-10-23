package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../service"
	"github.com/kataras/golog"
)

type ProductController struct {
	service.IProductService
}

func (controller *ProductController) GetAllProductHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("GET - Product: GetAllProductHandler (/products)")
	products, err := controller.GetAllProduct()
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func (controller *ProductController) GetProductsByNameHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("GET - Product: GetProductsByNameHandler (/products/search)")
	query := r.URL.Query().Get("query")
	products, err := controller.GetProductsByName(query)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}

func (controller *ProductController) UploadExcelProductHandler(w http.ResponseWriter, r *http.Request) {
	golog.Info("POST - Product: UploadExcelProductHandler (/products/upload_excel/:sheetName)")
	vars := mux.Vars(r)
	r.ParseMultipartForm(20 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()

	err = controller.NewProductUsingExcel(vars["sheetName"], file)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	products, err := controller.GetAllProduct()
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		golog.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
}
