package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/helper"
	"golang-rest-api/model/web"
	"golang-rest-api/service"
	"net/http"
	"strconv"
)

type OrdersControllerImpl struct {
	OrdersService service.OrdersService
}

func NewOrdersController(ordersService service.OrdersService) OrderController {
	return &OrdersControllerImpl{OrdersService: ordersService}
}

func (controller *OrdersControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersCreateRequest := web.OrdersCreateRequest{}
	helper.ReadFromRequestBody(request, &ordersCreateRequest)

	ordersResponse := controller.OrdersService.Create(request.Context(), ordersCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersUpdateRequest := web.OrdersUpdateRequest{}
	helper.ReadFromRequestBody(request, &ordersUpdateRequest)

	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	ordersUpdateRequest.Id = id

	ordersResponse := controller.OrdersService.Update(request.Context(), ordersUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	controller.OrdersService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	ordersResponse := controller.OrdersService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrdersControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersResponse := controller.OrdersService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
