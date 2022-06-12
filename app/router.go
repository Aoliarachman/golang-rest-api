package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/controller"
	"golang-rest-api/exception"
)

func NewRouter(categoryController controller.CategoryController, customerController controller.CustomerController, order_productController controller.OrderProductController, ordersController controller.OrderController, productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindByAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customerController.FindByAll)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.POST("/api/customers", customerController.Create)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)

	router.GET("/api/order_products", order_productController.FindByAll)
	router.GET("/api/order_products/:order_productId", order_productController.FindById)
	router.POST("/api/order_products", order_productController.Create)
	router.PUT("/api/order_products/:order_productId", order_productController.Update)
	router.DELETE("/api/order_products/:order_productId", order_productController.Delete)

	router.GET("/api/orders", ordersController.FindByAll)
	router.GET("/api/orders/:ordersId", ordersController.FindById)
	router.POST("/api/orders", ordersController.Create)
	router.PUT("/api/orders/:ordersId", ordersController.Update)
	router.DELETE("/api/orders/:ordersId", ordersController.Delete)

	router.GET("/api/product", productController.FindByAll)
	router.GET("/api/product/:productId", productController.FindById)
	router.POST("/api/product", productController.Create)
	router.PUT("/api/product/:productId", productController.Update)
	router.DELETE("/api/product/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
