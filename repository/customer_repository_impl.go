package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (c CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "insert into customer(name,address,email,phonenumber) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Address, customer.Email, customer.PhoneNumber)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer
}

func (c CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (c CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (c CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	SQL := "select id, name from customer where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("category is not found")
	}
}

func (c CustomerRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Customer {
	SQL := "select id, name from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers
}
