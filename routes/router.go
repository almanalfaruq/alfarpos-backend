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
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
	authMw := New(AuthConfig{
		SecretKey: config.SecretKey,
	})
	routesApi := routes.PathPrefix("/api/").Subrouter()
	routesApi.Use(CORS)

	userController := InjectUserController(database, config)
	userRoutes := routesApi.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/register", userController.RegisterHandler).Methods("POST", "OPTIONS")
	userRoutes.HandleFunc("/login", userController.LoginHandler).Methods("POST", "OPTIONS")
	userRoutes.HandleFunc("/update", userController.UpdateHandler).Methods("PUT", "OPTIONS")

	productController := InjectProductController(database, config)
	productRoutes := routesApi.PathPrefix("/products").Methods("GET").Subrouter()
	productRoutes.HandleFunc("", productController.GetProductsHandler).Methods("GET")
	productRoutes.HandleFunc("/id/{id}/", productController.GetProductByIdHandler).Methods("GET")
	productRoutes.HandleFunc("/ids/{ids}/", productController.GetProductsByIDsHandler).Methods("GET")
	productRoutes.HandleFunc("/code/{code}/", productController.GetProductByCodeHandler).Methods("GET")
	productRoutes.HandleFunc("/export_excel", productController.ExportAllProductsToExcelHandler).Methods("GET")

	productProtected := routesApi.PathPrefix("/products").Subrouter()
	productProtected.Use(authMw.JWT)
	productProtected.HandleFunc("", productController.NewProductHandler).Methods("POST", "OPTIONS")
	productProtected.HandleFunc("/upload_excel/{sheetName}", productController.UploadExcelProductHandler).Methods("POST")
	productProtected.HandleFunc("/{id}", productController.UpdateProductHandler).Methods("PUT", "OPTIONS")
	productProtected.HandleFunc("/upsert_excel/{sheetName}", productController.UpsertProductUsingExcelHandler).Methods("PUT")
	productProtected.HandleFunc("/{id}", productController.DeleteProductHandler).Methods("DELETE", "OPTIONS")

	categoryController := InjectCategoryController(database, config)
	categoryRoutes := routesApi.PathPrefix("/categories").Methods("GET").Subrouter()
	categoryRoutes.HandleFunc("", categoryController.GetCategoriesHandler).Methods("GET")
	categoryRoutes.HandleFunc("/id/{id}", categoryController.GetCategoryByIdHandler).Methods("GET")

	categoryProtected := routesApi.PathPrefix("/categories").Subrouter()
	categoryProtected.Use(authMw.JWT)
	categoryProtected.HandleFunc("", categoryController.NewCategoryHandler).Methods("POST")
	categoryProtected.HandleFunc("/{id}", categoryController.UpdateCategoryHandler).Methods("PUT")
	categoryProtected.HandleFunc("/{id}", categoryController.DeleteCategoryHandler).Methods("DELETE")

	unitController := InjectUnitController(database, config)
	unitRoutes := routesApi.PathPrefix("/units").Methods("GET").Subrouter()
	unitRoutes.HandleFunc("", unitController.GetUnitsHandler).Methods("GET")
	unitRoutes.HandleFunc("/id/{id}", unitController.GetUnitByIdHandler).Methods("GET")

	unitProtected := routesApi.PathPrefix("/units").Subrouter()
	unitProtected.Use(authMw.JWT)
	unitProtected.HandleFunc("", unitController.NewUnitHandler).Methods("POST")
	unitProtected.HandleFunc("/{id}", unitController.UpdateUnitHandler).Methods("PUT")
	unitProtected.HandleFunc("/{id}", unitController.DeleteUnitHandler).Methods("DELETE")

	paymentController := InjectPaymentController(database, config)
	paymentRoutes := routesApi.PathPrefix("/payments").Methods("GET").Subrouter()
	paymentRoutes.HandleFunc("", paymentController.GetPaymentsHandler).Methods("GET")
	paymentRoutes.HandleFunc("/id/{id}", paymentController.GetPaymentByIdHandler).Methods("GET")

	paymentProtected := routesApi.PathPrefix("/payments").Subrouter()
	paymentProtected.Use(authMw.JWT)
	paymentProtected.HandleFunc("", paymentController.NewPaymentHandler).Methods("POST")
	paymentProtected.HandleFunc("/{id}", paymentController.UpdatePaymentHandler).Methods("PUT")
	paymentProtected.HandleFunc("/{id}", paymentController.DeletePaymentHandler).Methods("DELETE")

	orderController := InjectOrderController(database, config)
	orderRoutes := routesApi.PathPrefix("/orders").Subrouter()
	orderRoutes.Use(authMw.JWT)
	orderRoutes.HandleFunc("", orderController.GetAllOrderHandler).Methods("GET")
	orderRoutes.HandleFunc("/id/{id}", orderController.GetOrderByIDHandler).Methods("GET")
	orderRoutes.HandleFunc("/filters", orderController.GetOrderUsingFilterHandler).Methods("GET", "POST")
	orderRoutes.HandleFunc("", orderController.NewOrderHandler).Methods("POST")
	orderRoutes.HandleFunc("/status", orderController.UpdateStatusHandler).Methods("PUT")

	printController := InjectPrintController(database, config)
	printRoutes := routesApi.PathPrefix("/print").Subrouter()
	printRoutes.Use(authMw.JWT)
	printRoutes.HandleFunc("/order/{invoice}", printController.OrderByInvoiceToPdfHandler).Methods("GET")

	moneyController := InjectMoneyController(database, config)
	moneyRoutes := routesApi.PathPrefix("/money").Subrouter()
	moneyRoutes.Use(authMw.JWT)
	moneyRoutes.HandleFunc("", moneyController.NewMoneyHandler).Methods("POST")
	moneyRoutes.HandleFunc("/filters", moneyController.GetMoneyTransactionWithFilterHandler).Methods("GET", "POST", "OPTIONS")

	profileController := InjectProfileController(database, config)
	profilesRoutes := routesApi.PathPrefix("/profiles").Subrouter()
	profilesRoutes.Use(authMw.JWT)
	profilesRoutes.HandleFunc("/{id}", profileController.GetProfileByIDHandler).Methods("GET")
	profilesRoutes.HandleFunc("", profileController.UpdateHandler).Methods("PUT")

	profileShopRoutes := routesApi.PathPrefix("/profile/shop").Subrouter()
	profileShopRoutes.Use(authMw.JWT)
	profileShopRoutes.HandleFunc("/", profileController.GetShopProfileHandler).Methods("GET")

	statsController := InjectStatsController(database, config)
	statsRoutes := routesApi.PathPrefix("/stats").Subrouter()
	statsRoutes.Use(authMw.JWT)
	statsRoutes.HandleFunc("", statsController.GetShopStats).Methods("GET", "OPTIONS")

	v2StatsRoutes := routesApi.PathPrefix("/v2/stats").Subrouter()
	v2StatsRoutes.Use(authMw.JWT)
	v2StatsRoutes.HandleFunc("/sales", statsController.GetShopStatsByDateRangeV2).Methods("GET", "OPTIONS")
	v2StatsRoutes.HandleFunc("/products/trending", statsController.GetTrendingProductsV2).Methods("GET", "OPTIONS")
	v2StatsRoutes.HandleFunc("/products/{id}/sales", statsController.GetProductSalesV2).Methods("GET", "OPTIONS")

	routes.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	return routes
}
