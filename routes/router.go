package routes

import (
	"net/http"

	. "github.com/almanalfaruq/alfarpos-backend/dependency_injection"
	_ "github.com/almanalfaruq/alfarpos-backend/docs"
	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetAllRoutes(database *util.DBConn, config util.Config) *mux.Router {
	routes := mux.NewRouter().StrictSlash(true)
	routes.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	authMw := New(AuthConfig{
		SecretKey: config.SecretKey,
	})
	routesApi := routes.PathPrefix("/api/").Subrouter()

	userController := InjectUserController(database, config)
	routesApi.HandleFunc("/users/register", userController.RegisterHandler).Methods("POST")
	routesApi.HandleFunc("/users/login", userController.LoginHandler).Methods("POST")

	productController := InjectProductController(database, config)
	routesApi.HandleFunc("/products", productController.GetProductsHandler).Methods("GET")
	routesApi.HandleFunc("/products/id/{id}/", productController.GetProductByIdHandler).Methods("GET")
	routesApi.HandleFunc("/products/code/{code}/", productController.GetProductByCodeHandler).Methods("GET")
	routesApi.HandleFunc("/products", authMw.CheckJWTToken(productController.NewProductHandler)).Methods("POST")
	routesApi.HandleFunc("/products/export_excel", productController.ExportAllProductsToExcelHandler).Methods("GET")
	routesApi.HandleFunc("/products/upload_excel/{sheetName}", authMw.CheckJWTToken(productController.UploadExcelProductHandler)).Methods("POST")
	routesApi.HandleFunc("/products/{id}", authMw.CheckJWTToken(productController.UpdateProductHandler)).Methods("PUT")

	categoryController := InjectCategoryController(database, config)
	routesApi.HandleFunc("/categories", categoryController.GetCategoriesHandler).Methods("GET")
	routesApi.HandleFunc("/categories/id/{id}", categoryController.GetCategoryByIdHandler).Methods("GET")
	routesApi.HandleFunc("/categories", authMw.CheckJWTToken(categoryController.NewCategoryHandler)).Methods("POST")
	routesApi.HandleFunc("/categories/{id}", authMw.CheckJWTToken(categoryController.UpdateCategoryHandler)).Methods("PUT")
	routesApi.HandleFunc("/categories/{id}", authMw.CheckJWTToken(categoryController.DeleteCategoryHandler)).Methods("DELETE")

	unitController := InjectUnitController(database, config)
	routesApi.HandleFunc("/units", unitController.GetUnitsHandler).Methods("GET")
	routesApi.HandleFunc("/units/id/{id}", unitController.GetUnitByIdHandler).Methods("GET")
	routesApi.HandleFunc("/units", authMw.CheckJWTToken(unitController.NewUnitHandler)).Methods("POST")
	routesApi.HandleFunc("/units/{id}", authMw.CheckJWTToken(unitController.UpdateUnitHandler)).Methods("PUT")
	routesApi.HandleFunc("/units/{id}", authMw.CheckJWTToken(unitController.DeleteUnitHandler)).Methods("DELETE")

	paymentController := InjectPaymentController(database, config)
	routesApi.HandleFunc("/payments", paymentController.GetPaymentsHandler).Methods("GET")
	routesApi.HandleFunc("/payments/id/{id}", paymentController.GetPaymentByIdHandler).Methods("GET")
	routesApi.HandleFunc("/payments", authMw.CheckJWTToken(paymentController.NewPaymentHandler)).Methods("POST")
	routesApi.HandleFunc("/payments/{id}", authMw.CheckJWTToken(paymentController.UpdatePaymentHandler)).Methods("PUT")
	routesApi.HandleFunc("/payments/{id}", authMw.CheckJWTToken(paymentController.DeletePaymentHandler)).Methods("DELETE")

	orderController := InjectOrderController(database, config)
	routesApi.HandleFunc("/orders", authMw.CheckJWTToken(orderController.GetAllOrderHandler)).Methods("GET")
	routesApi.HandleFunc("/orders/id/{id}", authMw.CheckJWTToken(orderController.GetOrderByIDHandler)).Methods("GET")
	routesApi.HandleFunc("/orders", orderController.NewOrderHandler).Methods("POST")

	printController := InjectPrintController(database, config)
	routesApi.HandleFunc("/print/order/{invoice}", authMw.CheckJWTToken(printController.OrderByInvoiceToPdfHandler)).Methods("GET")

	routes.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	return routes
}
