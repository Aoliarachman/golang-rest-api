package service

import (
	"context"
	"golang-rest-api/model/web"
)

type OrderProductService interface {
	Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse
	Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse
	Delete(ctx context.Context, order_productId int)
	FindById(ctx context.Context, order_productId int) web.OrderProductResponse
	FindByAll(ctx context.Context) []web.OrderProductResponse
}
