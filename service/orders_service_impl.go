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

type OrdersServiceImpl struct {
	OrdersRepository repository.OrdersRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrdersService(ordersRepository repository.OrdersRepository, DB *sql.DB, validate *validator.Validate) OrdersService {

	return &OrdersServiceImpl{
		OrdersRepository: ordersRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *OrdersServiceImpl) Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := domain.Orders{
		Id: request.Id,
	}
	orders = service.OrdersRepository.Save(ctx, tx, orders)

	return helper.ToOrdersResponse(orders)
}

func (service OrdersServiceImpl) Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders, err := service.OrdersRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	orders.Id = request.Id
	orders = service.OrdersRepository.Update(ctx, tx, orders)

	return helper.ToOrdersResponse(orders)
}

func (service OrdersServiceImpl) Delete(ctx context.Context, ordersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders, err := service.OrdersRepository.FindById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrdersRepository.Delete(ctx, tx, orders)
}

func (service OrdersServiceImpl) FindById(ctx context.Context, ordersId int) web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders, err := service.OrdersRepository.FindById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrdersResponse(orders)
}

func (service OrdersServiceImpl) FindByAll(ctx context.Context) []web.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrdersRepository.FindByAll(ctx, tx)

	return helper.ToOrdersResponses(orders)

}
