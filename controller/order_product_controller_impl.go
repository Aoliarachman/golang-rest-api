package controller

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/helper"
	"golang-rest-api/model/web"
	"golang-rest-api/service"
	"net/http"
	"strconv"
)

type OrderProductControllerImpl struct {
	OrderProductService service.OrderProductService
}

func NewOrderProductController(order_productService service.OrderProductService) OrderProductController {
	return &OrderProductControllerImpl{OrderProductService: order_productService}
}

func (controller *OrderProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	order_productCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &order_productCreateRequest)

	order_productResponse := controller.OrderProductService.Create(request.Context(), order_productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   order_productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	OrderProductUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &OrderProductUpdateRequest)

	order_productId := params.ByName("order_productId")
	id, err := strconv.Atoi(order_productId)
	helper.PanicIfError(err)

	OrderProductUpdateRequest.Id = id

	order_productResponse := controller.OrderProductService.Update(request.Context(), OrderProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   order_productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	order_productId := params.ByName("order_productId")
	id, err := strconv.Atoi(order_productId)
	helper.PanicIfError(err)

	controller.OrderProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	order_productId := params.ByName("order_productId")
	id, err := strconv.Atoi(order_productId)
	helper.PanicIfError(err)

	order_productResponse := controller.OrderProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   order_productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	order_productResponse := controller.OrderProductService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   order_productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
