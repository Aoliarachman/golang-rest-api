package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

type OrderProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct) domain.OrderProduct
	Update(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct) domain.OrderProduct
	Delete(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct)
	FindById(ctx context.Context, tx *sql.Tx, order_productId int) (domain.OrderProduct, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct
}
