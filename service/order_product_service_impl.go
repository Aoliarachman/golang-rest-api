package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-rest-api/exception"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
	"golang-rest-api/repository"
)

type OrderProductServiceImpl struct {
	OrderProductRepository repository.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(order_productRepository repository.OrderProductRepository, DB *sql.DB, validate *validator.Validate) *OrderProductServiceImpl {

	return &OrderProductServiceImpl{
		OrderProductRepository: order_productRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *OrderProductServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order_product := domain.OrderProduct{
		Id:        request.Id,
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Qty:       request.Qty,
		Price:     request.Price,
		Amount:    request.Amount,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
	order_product = service.OrderProductRepository.Save(ctx, tx, order_product)

	return helper.ToOrderProductResponse(order_product)
}

func (service *OrderProductServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order_product, err := service.OrderProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	order_product.Id = request.Id
	order_product.OrderId = request.OrderId
	order_product.ProductId = request.ProductId
	order_product.Qty = request.Qty
	order_product.Price = request.Price
	order_product.Amount = request.Amount

	order_product = service.OrderProductRepository.Update(ctx, tx, order_product)

	return helper.ToOrderProductResponse(order_product)
}

func (service *OrderProductServiceImpl) Delete(ctx context.Context, order_productId int) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order_product, err := service.OrderProductRepository.FindById(ctx, tx, order_productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrderProductRepository.Delete(ctx, tx, order_product)
}

func (service *OrderProductServiceImpl) FindById(ctx context.Context, order_productId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order_product, err := service.OrderProductRepository.FindById(ctx, tx, order_productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrderProductResponse(order_product)
}

func (service OrderProductServiceImpl) FindByAll(ctx context.Context) []web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order_product := service.OrderProductRepository.FindByAll(ctx, tx)

	return helper.ToOrderProductResponses(order_product)
}
