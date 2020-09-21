package routes

import (
	"net/http"

	. "github.com/almanalfaruq/alfarpos-backend/dependency_injection"
	_ "github.com/almanalfaruq/alfarpos-backend/docs"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/almanalfaruq/alfarpos-backend/util/auth"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetAllRoutes(database *util.DBConn, config util.Config) *mux.Router {
	mw := auth.New(config)
	routes := mux.NewRouter().StrictSlash(true)
	routes.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("https://alfarpos-backend.herokuapp.com/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	routesApi := routes.PathPrefix("/api/").Subrouter()

	userController := InjectUserController(database, config)
	routesApi.HandleFunc("/users/register", userController.RegisterHandler).Methods("POST")
	routesApi.HandleFunc("/users/login", userController.LoginHandler).Methods("POST")

	productController := InjectProductController(database, config)
	routesApi.HandleFunc("/products", productController.GetProductsHandler).Methods("GET")
	routesApi.HandleFunc("/products/id/{id}/", productController.GetProductByIdHandler).Methods("GET")
	routesApi.HandleFunc("/products/code/{code}/", productController.GetProductByCodeHandler).Methods("GET")
	routesApi.HandleFunc("/products", mw.CheckAuth(productController.NewProductHandler)).Methods("POST")
	routesApi.HandleFunc("/products/export_excel", productController.ExportAllProductsToExcelHandler).Methods("GET")
	routesApi.HandleFunc("/products/upload_excel/{sheetName}", mw.CheckAuth(productController.UploadExcelProductHandler)).Methods("POST")
	routesApi.HandleFunc("/products/{id}", mw.CheckAuth(productController.UpdateProductHandler)).Methods("PUT")

	categoryController := InjectCategoryController(database, config)
	routesApi.HandleFunc("/categories", categoryController.GetCategoriesHandler).Methods("GET")
	routesApi.HandleFunc("/categories/id/{id}", categoryController.GetCategoryByIdHandler).Methods("GET")
	routesApi.HandleFunc("/categories", mw.CheckAuth(categoryController.NewCategoryHandler)).Methods("POST")
	routesApi.HandleFunc("/categories/{id}", mw.CheckAuth(categoryController.UpdateCategoryHandler)).Methods("PUT")
	routesApi.HandleFunc("/categories/{id}", mw.CheckAuth(categoryController.DeleteCategoryHandler)).Methods("DELETE")

	unitController := InjectUnitController(database, config)
	routesApi.HandleFunc("/units", unitController.GetUnitsHandler).Methods("GET")
	routesApi.HandleFunc("/units/id/{id}", unitController.GetUnitByIdHandler).Methods("GET")
	routesApi.HandleFunc("/units", mw.CheckAuth(unitController.NewUnitHandler)).Methods("POST")
	routesApi.HandleFunc("/units/{id}", mw.CheckAuth(unitController.UpdateUnitHandler)).Methods("PUT")
	routesApi.HandleFunc("/units/{id}", mw.CheckAuth(unitController.DeleteUnitHandler)).Methods("DELETE")

	paymentController := InjectPaymentController(database, config)
	routesApi.HandleFunc("/payments", paymentController.GetPaymentsHandler).Methods("GET")
	routesApi.HandleFunc("/payments/id/{id}", paymentController.GetPaymentByIdHandler).Methods("GET")
	routesApi.HandleFunc("/payments", mw.CheckAuth(paymentController.NewPaymentHandler)).Methods("POST")
	routesApi.HandleFunc("/payments/{id}", mw.CheckAuth(paymentController.UpdatePaymentHandler)).Methods("PUT")
	routesApi.HandleFunc("/payments/{id}", mw.CheckAuth(paymentController.DeletePaymentHandler)).Methods("DELETE")

	orderController := InjectOrderController(database, config)
	routesApi.HandleFunc("/orders", mw.CheckAuth(orderController.GetAllOrderHandler)).Methods("GET")
	routesApi.HandleFunc("/orders", mw.CheckAuth(orderController.NewOrderHandler)).Methods("POST")

	printController := InjectPrintController(database, config)
	routesApi.HandleFunc("/print/order/{invoice}", mw.CheckAuth(printController.OrderByInvoiceToPdfHandler)).Methods("GET")

	routes.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	return routes
}
