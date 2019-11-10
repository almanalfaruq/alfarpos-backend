package dependency_injection

import (
	"github.com/almanalfaruq/alfarpos-backend/controller"
	"github.com/almanalfaruq/alfarpos-backend/repository"
	"github.com/almanalfaruq/alfarpos-backend/service"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

func InjectUserController(databaseConnection *util.DatabaseConnection, config util.Config) controller.UserController {
	userRepository := &repository.UserRepository{
		IDatabaseConnection: databaseConnection,
	}
	userService := &service.UserService{
		User:   userRepository,
		Config: config,
	}
	userController := controller.UserController{
		IUserService: userService,
	}
	return userController
}

func InjectProductController(databaseConnection *util.DatabaseConnection, config util.Config) controller.ProductController {
	productRepository := &repository.ProductRepository{
		IDatabaseConnection: databaseConnection,
	}
	categoryRepository := &repository.CategoryRepository{
		IDatabaseConnection: databaseConnection,
	}
	unitRepository := &repository.UnitRepository{
		IDatabaseConnection: databaseConnection,
	}
	stockRepository := &repository.StockRepository{
		IDatabaseConnection: databaseConnection,
	}
	productService := &service.ProductService{
		Product:  productRepository,
		Category: categoryRepository,
		Unit:     unitRepository,
		Stock:    stockRepository,
	}
	productController := controller.ProductController{
		IProductService: productService,
		Config:          config,
	}
	return productController
}

func InjectOrderController(databaseConnection *util.DatabaseConnection, config util.Config) controller.OrderController {
	orderRepository := &repository.OrderRepository{
		IDatabaseConnection: databaseConnection,
	}
	orderDetailRepository := &repository.OrderDetailRepository{
		IDatabaseConnection: databaseConnection,
	}
	paymentRepository := &repository.PaymentRepository{
		IDatabaseConnection: databaseConnection,
	}
	customerRepository := &repository.CustomerRepository{
		IDatabaseConnection: databaseConnection,
	}
	productRepository := &repository.ProductRepository{
		IDatabaseConnection: databaseConnection,
	}
	orderService := &service.OrderService{
		Order:       orderRepository,
		OrderDetail: orderDetailRepository,
		Payment:     paymentRepository,
		Customer:    customerRepository,
		Product:     productRepository,
	}
	orderController := controller.OrderController{
		IOrderService: orderService,
		Config:        config,
	}
	return orderController
}

func InjectCategoryController(databaseConnection *util.DatabaseConnection, config util.Config) controller.CategoryController {
	categoryRepository := &repository.CategoryRepository{
		IDatabaseConnection: databaseConnection,
	}
	categoryService := &service.CategoryService{
		ICategoryRepository: categoryRepository,
	}
	categoryController := controller.CategoryController{
		ICategoryService: categoryService,
		Config:           config,
	}
	return categoryController
}

func InjectUnitController(databaseConnection *util.DatabaseConnection, config util.Config) controller.UnitController {
	unitRepository := &repository.UnitRepository{
		IDatabaseConnection: databaseConnection,
	}
	unitService := &service.UnitService{
		IUnitRepository: unitRepository,
	}
	unitController := controller.UnitController{
		IUnitService: unitService,
		Config:       config,
	}
	return unitController
}

func InjectPaymentController(databaseConnection *util.DatabaseConnection, config util.Config) controller.PaymentController {
	paymentRepository := &repository.PaymentRepository{
		IDatabaseConnection: databaseConnection,
	}
	paymentService := &service.PaymentService{
		IPaymentRepository: paymentRepository,
	}
	paymentController := controller.PaymentController{
		IPaymentService: paymentService,
		Config:          config,
	}
	return paymentController
}

func InjectPrintController(databaseConnection *util.DatabaseConnection, config util.Config) controller.PrintController {
	orderRepository := &repository.OrderRepository{
		IDatabaseConnection: databaseConnection,
	}
	printService := &service.PrintService{
		Order:  orderRepository,
		Config: config,
	}
	printController := controller.PrintController{
		IPrintService: printService,
	}
	return printController
}
