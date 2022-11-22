package web

import "time"

type OrderCreateRequest struct {
	CustomerName string    `json:"customer_name" validate:"required"`
	OrderedAt    time.Time `json:"ordered_at"`
}

type OrderUpdateRequest struct {
	OrderID      int       `json:"order_id" validate:"required"`
	CustomerName string    `json:"customer_name" validate:"required"`
	OrderedAt    time.Time `json:"ordered_at"`
}

type OrderResponse struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}
