package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	order_productRepository := repository.NewOrderProductRepository()
	order_productService := service.NewOrderProductService(order_productRepository, db, validate)
	order_productController := controller.NewOrderProductController(order_productService)

	ordersRepository := repository.NewOrdersRepository()
	ordersService := service.NewOrdersService(ordersRepository, db, validate)
	ordersController := controller.NewOrdersController(ordersService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(categoryController, customerController, order_productController, productController, ordersController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
