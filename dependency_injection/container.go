package dependency_injection

import (
	"../controller"
	"../repository"
	"../service"
	"../util"
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

func InjectProductController(databaseConnection *util.DatabaseConnection) controller.ProductController {
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
	}
	return productController
}
