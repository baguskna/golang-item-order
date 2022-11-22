package service

import (
	"context"
	"database/sql"
	"golang-item-order/helper"
	"golang-item-order/model/domain"
	"golang-item-order/model/web"
	"golang-item-order/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &CategoryServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (s *CategoryServiceImpl) Save(c context.Context, r web.OrderCreateRequest) web.OrderResponse {
	err := s.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		CustomerName: r.CustomerName,
		OrderedAt:    r.OrderedAt,
	}

	order = s.OrderRepository.Save(c, tx, order)

	return helper.ToCategoryResponse(order)
}

func (s *CategoryServiceImpl) Update(c context.Context, r web.OrderUpdateRequest) web.OrderResponse {
	err := s.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := s.OrderRepository.FindById(c, tx, r.OrderID)

	order.CustomerName = r.CustomerName

	order = s.OrderRepository.Update(c, tx, order)

	return helper.ToCategoryResponse(order)
}

func (s *CategoryServiceImpl) Delete(c context.Context, orderId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := s.OrderRepository.FindById(c, tx, orderId)
	helper.PanicIfError(err)

	s.OrderRepository.Delete(c, tx, order)
}

func (s *CategoryServiceImpl) FindById(c context.Context, orderId int) web.OrderResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := s.OrderRepository.FindById(c, tx, orderId)
	helper.PanicIfError(err)

	return web.OrderResponse(order)
}

func (s *CategoryServiceImpl) FindAll(c context.Context) []web.OrderResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := s.OrderRepository.FindAll(c, tx)

	var ordersResponse []web.OrderResponse
	for _, order := range orders {
		ordersResponse = append(ordersResponse, web.OrderResponse(order))
	}
	return ordersResponse
}
