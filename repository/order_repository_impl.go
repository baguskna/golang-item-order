package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-item-order/helper"
	"golang-item-order/model/domain"
	"time"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) Save(c context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "insert into orders(customer_name, ordered_at) value(?, ?)"
	result, err := tx.ExecContext(c, SQL, order.CustomerName, time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.OrderID = int(id)
	return order
}

func (r *OrderRepositoryImpl) Update(c context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	SQL := "update orders set customer_name = ? where id = ?"
	_, err := tx.ExecContext(c, SQL, order.CustomerName, order.OrderID)
	helper.PanicIfError(err)

	return order
}

func (r *OrderRepositoryImpl) Delete(c context.Context, tx *sql.Tx, order domain.Order) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(c, SQL, order.CustomerName, order.OrderID)
	helper.PanicIfError(err)
}

func (r *OrderRepositoryImpl) FindById(c context.Context, tx *sql.Tx, categoryId int) (domain.Order, error) {
	SQL := "select order_id, customer_name, ordered_at from orders where id = ?"
	rows, err := tx.QueryContext(c, SQL, categoryId)
	helper.PanicIfError(err)

	order := domain.Order{}
	// if data exists
	if rows.Next() {
		err := rows.Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, errors.New("Category is not found")
	}
}

func (r *OrderRepositoryImpl) FindAll(c context.Context, tx *sql.Tx) []domain.Order {
	SQL := "select order_id, customer_name, ordered_at from orders"
	rows, err := tx.QueryContext(c, SQL)
	helper.PanicIfError(err)

	var orders []domain.Order
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}

	return orders
}
