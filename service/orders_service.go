package service

import (
	"context"
	"golang-rest-api/model/web"
)

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse
	Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse
	Delete(ctx context.Context, order_productId int)
	FindById(ctx context.Context, order_productId int) web.OrdersResponse
	FindByAll(ctx context.Context) []web.OrdersResponse
}
