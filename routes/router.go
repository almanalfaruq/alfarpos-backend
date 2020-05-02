package routes

import (
	. "github.com/almanalfaruq/alfarpos-backend/dependency_injection"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/gorilla/mux"
)

func GetAllRoutes(database *util.DBConn, config util.Config) *mux.Router {
	routes := mux.NewRouter().StrictSlash(true).PathPrefix("/api/").Subrouter()

	userController := InjectUserController(database, config)
	routes.HandleFunc("/users/register", userController.RegisterHandler).Methods("POST")
	routes.HandleFunc("/users/login", userController.LoginHandler).Methods("POST")

	productController := InjectProductController(database, config)
	routes.HandleFunc("/products", productController.GetProductsHandler).Methods("GET")
	routes.HandleFunc("/products/id/{id}/", productController.GetProductByIdHandler).Methods("GET")
	routes.HandleFunc("/products/code/{code}/", productController.GetProductByCodeHandler).Methods("GET")
	routes.HandleFunc("/products", productController.NewProductHandler).Methods("POST")
	routes.HandleFunc("/products/export_excel", productController.ExportAllProductsToExcelHandler).Methods("GET")
	routes.HandleFunc("/products/upload_excel/{sheetName}", productController.UploadExcelProductHandler).Methods("POST")
	routes.HandleFunc("/products/{id}", productController.UpdateProductHandler).Methods("PUT")

	categoryController := InjectCategoryController(database, config)
	routes.HandleFunc("/categories", categoryController.GetCategoriesHandler).Methods("GET")
	routes.HandleFunc("/categories/id/{id}", categoryController.GetCategoryByIdHandler).Methods("GET")
	routes.HandleFunc("/categories", categoryController.NewCategoryHandler).Methods("POST")
	routes.HandleFunc("/categories/{id}", categoryController.UpdateCategoryHandler).Methods("PUT")
	routes.HandleFunc("/categories/{id}", categoryController.DeleteCategoryHandler).Methods("DELETE")

	unitController := InjectUnitController(database, config)
	routes.HandleFunc("/units", unitController.GetUnitsHandler).Methods("GET")
	routes.HandleFunc("/units/id/{id}", unitController.GetUnitByIdHandler).Methods("GET")
	routes.HandleFunc("/units", unitController.NewUnitHandler).Methods("POST")
	routes.HandleFunc("/units/{id}", unitController.UpdateUnitHandler).Methods("PUT")
	routes.HandleFunc("/units/{id}", unitController.DeleteUnitHandler).Methods("DELETE")

	paymentController := InjectPaymentController(database, config)
	routes.HandleFunc("/payments", paymentController.GetPaymentsHandler).Methods("GET")
	routes.HandleFunc("/payments/id/{id}", paymentController.GetPaymentByIdHandler).Methods("GET")
	routes.HandleFunc("/payments", paymentController.NewPaymentHandler).Methods("POST")
	routes.HandleFunc("/payments/{id}", paymentController.UpdatePaymentHandler).Methods("PUT")
	routes.HandleFunc("/payments/{id}", paymentController.DeletePaymentHandler).Methods("DELETE")

	orderController := InjectOrderController(database, config)
	routes.HandleFunc("/orders", orderController.GetAllOrderHandler).Methods("GET")
	routes.HandleFunc("/orders", orderController.NewOrderHandler).Methods("POST")

	printController := InjectPrintController(database, config)
	routes.HandleFunc("/print/order/{invoice}", printController.OrderByInvoiceToPdfHandler).Methods("GET")
	return routes
}
