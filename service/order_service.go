package service

import (
	"context"
	"golang-item-order/model/web"
)

type OrderService interface {
	Save(c context.Context, r web.OrderCreateRequest) web.OrderResponse
	Update(c context.Context, r web.OrderUpdateRequest) web.OrderResponse
	Delete(c context.Context, orderId int)
	FindById(c context.Context, orderId int) web.OrderResponse
	FindAll(c context.Context) []web.OrderResponse
}
