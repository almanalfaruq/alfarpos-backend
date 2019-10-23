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

	productController := InjectProductController(database)
	routes.HandleFunc("/products", productController.GetAllProductHandler).Methods("GET")
	routes.HandleFunc("/products/search", productController.GetProductsByNameHandler).Methods("GET")
	routes.HandleFunc("/products/upload_excel/{sheetName}", productController.UploadExcelProductHandler).Methods("POST")
	return routes
}
