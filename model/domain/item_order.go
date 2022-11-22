package domain

import "time"

type Item struct {
	ItemID      int
	ItemCode    int
	Description string
	Quantity    int
	Order       []Order
}

type Order struct {
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
}
