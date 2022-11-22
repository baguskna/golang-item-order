package main

import (
	"golang-item-order/app"
	"golang-item-order/controller"
	"golang-item-order/helper"
	"golang-item-order/repository"
	"golang-item-order/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	r := httprouter.New()

	r.GET("/api/orders", orderController.FindAll)
	r.GET("/api/orders/:orderId", orderController.FindById)
	r.POST("/api/orders", orderController.Create)
	r.PUT("/api/orders/:orderId", orderController.Update)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
