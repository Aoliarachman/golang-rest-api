package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (p ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id, product.Id)
	helper.PanicIfError(err)

	return product
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "select id, Name from product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (p ProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select id, Name from product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}
