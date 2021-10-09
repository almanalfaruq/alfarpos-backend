package dependency_injection

import (
	"github.com/almanalfaruq/alfarpos-backend/controller"
	orderctrl "github.com/almanalfaruq/alfarpos-backend/controller/order"
	productctrl "github.com/almanalfaruq/alfarpos-backend/controller/product"
	profilectrl "github.com/almanalfaruq/alfarpos-backend/controller/profile"
	statsctrl "github.com/almanalfaruq/alfarpos-backend/controller/stats"
	transactionctrl "github.com/almanalfaruq/alfarpos-backend/controller/transaction"
	userctrl "github.com/almanalfaruq/alfarpos-backend/controller/user"
	"github.com/almanalfaruq/alfarpos-backend/repository"
	orderrepo "github.com/almanalfaruq/alfarpos-backend/repository/order"
	productrepo "github.com/almanalfaruq/alfarpos-backend/repository/product"
	profilerepo "github.com/almanalfaruq/alfarpos-backend/repository/profile"
	statsrepo "github.com/almanalfaruq/alfarpos-backend/repository/stats"
	stockrepo "github.com/almanalfaruq/alfarpos-backend/repository/stock"
	transactionrepo "github.com/almanalfaruq/alfarpos-backend/repository/transaction"
	userrepo "github.com/almanalfaruq/alfarpos-backend/repository/user"
	"github.com/almanalfaruq/alfarpos-backend/service"
	ordersvc "github.com/almanalfaruq/alfarpos-backend/service/order"
	productsvc "github.com/almanalfaruq/alfarpos-backend/service/product"
	profilesvc "github.com/almanalfaruq/alfarpos-backend/service/profile"
	statssvc "github.com/almanalfaruq/alfarpos-backend/service/stats"
	transactionsvc "github.com/almanalfaruq/alfarpos-backend/service/transaction"
	usersvc "github.com/almanalfaruq/alfarpos-backend/service/user"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

func InjectUserController(dbConn *util.DBConn, config util.Config) *userctrl.UserController {
	userRepo := userrepo.NewUserRepo(dbConn)
	userService := usersvc.NewUserService(config, userRepo)
	return userctrl.NewUserController(userService)
}

func InjectProductController(dbConn *util.DBConn, config util.Config) *productctrl.ProductController {
	productRepository := productrepo.NewProductRepo(dbConn)
	categoryRepository := repository.NewCategoryRepo(dbConn)
	unitRepository := repository.NewUnitRepo(dbConn)
	stockRepository := stockrepo.NewStockRepo(dbConn)
	productService := productsvc.NewProductService(productRepository, categoryRepository, unitRepository, stockRepository)
	productController := productctrl.NewProductController(config, productService)
	return productController
}

func InjectOrderController(dbConn *util.DBConn, config util.Config) *orderctrl.OrderController {
	orderRepository := orderrepo.NewOrderRepo(dbConn)
	orderDetailRepository := orderrepo.NewOrderDetailRepo(dbConn)
	paymentRepository := repository.NewPaymentRepo(dbConn)
	customerRepository := repository.NewCustomerRepo(dbConn)
	productRepository := productrepo.NewProductRepo(dbConn)
	orderService := ordersvc.NewOrderService(orderRepository, orderDetailRepository, paymentRepository, customerRepository,
		productRepository)
	orderController := orderctrl.NewOrderController(config, orderService)
	return orderController
}

func InjectCategoryController(dbConn *util.DBConn, config util.Config) *controller.CategoryController {
	categoryRepository := repository.NewCategoryRepo(dbConn)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(config, categoryService)
	return categoryController
}

func InjectUnitController(dbConn *util.DBConn, config util.Config) *controller.UnitController {
	unitRepository := repository.NewUnitRepo(dbConn)
	unitService := service.NewUnitService(unitRepository)
	unitController := controller.NewUnitController(config, unitService)
	return unitController
}

func InjectPaymentController(dbConn *util.DBConn, config util.Config) *controller.PaymentController {
	paymentRepository := repository.NewPaymentRepo(dbConn)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentController := controller.NewPaymentController(config, paymentService)
	return paymentController
}

func InjectPrintController(dbConn *util.DBConn, config util.Config) *controller.PrintController {
	orderRepository := orderrepo.NewOrderRepo(dbConn)
	printService := service.NewPrintService(config, orderRepository)
	printController := controller.NewPrintController(printService)
	return printController
}

func InjectMoneyController(dbConn *util.DBConn, config util.Config) *transactionctrl.MoneyController {
	moneyRepo := transactionrepo.NewMoney(dbConn)
	moneyService := transactionsvc.New(moneyRepo)
	return transactionctrl.NewMoney(moneyService)
}

func InjectProfileController(dbConn *util.DBConn, config util.Config) *profilectrl.ProfileController {
	profileRepo := profilerepo.NewProfile(dbConn)
	profileService := profilesvc.NewProfile(&config, profileRepo)
	return profilectrl.NewProfile(profileService)
}

func InjectStatsController(dbConn *util.DBConn, config util.Config) *statsctrl.StatsController {
	statsRepo := statsrepo.NewStats(dbConn)
	orderRepo := orderrepo.NewOrderRepo(dbConn)
	moneyRepo := transactionrepo.NewMoney(dbConn)
	statsService := statssvc.New(statsRepo, orderRepo, moneyRepo)
	return statsctrl.New(statsService)
}
