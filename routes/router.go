package routes

import (
	. "../dependency_injection"
	"../util"
	"github.com/gorilla/mux"
)

func GetAllRoutes(database *util.DatabaseConnection, config util.Config) *mux.Router {
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	userController := InjectUserController(database, config)
	routes.HandleFunc("/users/register", userController.RegisterHandler).Methods("POST")
	routes.HandleFunc("/users/login", userController.LoginHandler).Methods("POST")

	productController := InjectProductController(database, config)
	routes.HandleFunc("/products", productController.GetProductsHandler).Methods("GET")
	routes.HandleFunc("/products/id/{id}/", productController.GetProductByIdHandler).Methods("GET")
	routes.HandleFunc("/products/code/{code}/", productController.GetProductByCodeHandler).Methods("GET")
	routes.HandleFunc("/products", productController.NewProductHandler).Methods("POST")
	routes.HandleFunc("/products/upload_excel/{sheetName}", productController.UploadExcelProductHandler).Methods("POST")
	routes.HandleFunc("/products/{id}", productController.UpdateProductHandler).Methods("PUT")
	return routes
}
