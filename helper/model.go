package helper

import (
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		Address:     customer.Address,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
	}
}

func ToCustomerResponses(customer []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customer {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}

func ToOrderProductResponse(order_product domain.OrderProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:        order_product.Id,
		OrderId:   order_product.OrderId,
		ProductId: order_product.ProductId,
		Qty:       order_product.Qty,
		Price:     order_product.Price,
		Amount:    order_product.Amount,
	}
}

func ToOrderProductResponses(order_product []domain.OrderProduct) []web.OrderProductResponse {
	var order_productResponses []web.OrderProductResponse
	for _, order_product := range order_product {
		order_productResponses = append(order_productResponses, ToOrderProductResponse(order_product))
	}
	return order_productResponses
}

func ToOrdersResponse(orders domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id:          orders.Id,
		OrderDate:   orders.OrderDate,
		CustomerId:  orders.CustomerId,
		TotalAmount: orders.TotalAmount,
	}
}

func ToOrdersResponses(orders []domain.Orders) []web.OrdersResponse {
	var ordersResponses []web.OrdersResponse
	for _, orders := range orders {
		ordersResponses = append(ordersResponses, ToOrdersResponse(orders))
	}
	return ordersResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToProductResponses(product []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range product {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
