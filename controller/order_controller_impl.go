package controller

import (
	"encoding/json"
	"golang-item-order/helper"
	"golang-item-order/model/web"
	"golang-item-order/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (c *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	orderCreateRequest := web.OrderCreateRequest{}
	err := decoder.Decode(&orderCreateRequest)
	helper.PanicIfError(err)

	orderResponse := c.OrderService.Save(r.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (c *OrderControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	orderUpdateRequest := web.OrderUpdateRequest{}
	err := decoder.Decode(&orderUpdateRequest)
	helper.PanicIfError(err)

	categoryId := params.ByName("orderId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	orderUpdateRequest.OrderID = id

	// categoryResponse := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
	orderResponse := c.OrderService.Update(r.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (c *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	c.OrderService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (c *OrderControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := c.OrderService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (c *OrderControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ordersResponse := c.OrderService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
