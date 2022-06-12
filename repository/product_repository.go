package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orders domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, orders domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, orders domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Product, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
