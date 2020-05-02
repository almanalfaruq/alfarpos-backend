package dependency_injection

import (
	"github.com/almanalfaruq/alfarpos-backend/controller"
	"github.com/almanalfaruq/alfarpos-backend/repository"
	"github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

func InjectUserController(dbConn *util.DBConn, config util.Config) *controller.UserController {
	userRepo := repository.NewUserRepo(dbConn)
	userService := service.NewUserService(config, userRepo)
	userController := controller.NewUserController(userService)
	return userController
}

func InjectProductController(dbConn *util.DBConn, config util.Config) *controller.ProductController {
	productRepository := repository.NewProductRepo(dbConn)
	categoryRepository := repository.NewCategoryRepo(dbConn)
	unitRepository := repository.NewUnitRepo(dbConn)
	stockRepository := repository.NewStockRepo(dbConn)
	productService := service.NewProductService(productRepository, categoryRepository, unitRepository, stockRepository)
	productController := controller.NewProductController(config, productService)
	return productController
}

func InjectOrderController(dbConn *util.DBConn, config util.Config) *controller.OrderController {
	orderRepository := repository.NewOrderRepo(dbConn)
	orderDetailRepository := repository.NewOrderDetailRepo(dbConn)
	paymentRepository := repository.NewPaymentRepo(dbConn)
	customerRepository := repository.NewCustomerRepo(dbConn)
	productRepository := repository.NewProductRepo(dbConn)
	orderService := service.NewOrderService(orderRepository, orderDetailRepository, paymentRepository, customerRepository,
		productRepository)
	orderController := controller.NewOrderController(config, orderService)
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
	orderRepository := repository.NewOrderRepo(dbConn)
	printService := service.NewPrintService(config, orderRepository)
	printController := controller.NewPrintController(printService)
	return printController
}
