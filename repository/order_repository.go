package repository

import (
	"context"
	"database/sql"
	"golang-item-order/model/domain"
)

type OrderRepository interface {
	Save(c context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Update(c context.Context, tx *sql.Tx, order domain.Order) domain.Order
	Delete(c context.Context, tx *sql.Tx, order domain.Order)
	FindById(c context.Context, tx *sql.Tx, categoryId int) (domain.Order, error)
	FindAll(c context.Context, tx *sql.Tx) []domain.Order
}
