package helper

import (
	"golang-item-order/model/domain"
	"golang-item-order/model/web"
)

func ToCategoryResponse(r domain.Order) web.OrderResponse {
	return web.OrderResponse{
		OrderID:      r.OrderID,
		CustomerName: r.CustomerName,
		OrderedAt:    r.OrderedAt,
	}
}
