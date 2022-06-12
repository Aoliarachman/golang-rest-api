package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/model/domain"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDBcustomer() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_rest_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRoutercustomer(db *sql.DB) http.Handler {
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)
	router := app.NewRouter(customerController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCustomer(db *sql.DB) {
	db.Exec("TRUNCATE customer")
}

func TestCreateCustomer(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCustomerFailed(t *testing.T) {

	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestUpdateCustomerSucces(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRespository := repository.NewCategoryRepository()
	category := categoryRespository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categpries"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestUpdateCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRespository := repository.NewCategoryRepository()
	category := categoryRespository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categpries"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])

}

func TestGetCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRespository := repository.NewCategoryRepository()
	category := categoryRespository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categpries/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])

}

func TestGetCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categpries/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}

func TestDeleteCustomerSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRespository := repository.NewCategoryRepository()
	category := categoryRespository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categpries/"+strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

}

func TestDeleteCustomerFailed(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/customer/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])

}

func TestListCustomersSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	customerRespository := repository.NewCustomerRepository()
	customer1 := customerRespository.Save(context.Background(), tx, domain.Customer{
		Name: "Gadget",
	})
	customer2 := customerRespository.Save(context.Background(), tx, domain.Customer{
		Name: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categpries", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])

	fmt.Println(responseBody)

	var customers = responseBody["data"].([]interface{})

	customerResponse1 := customers[0].(map[string]interface{})
	customerResponse2 := customers[1].(map[string]interface{})

	assert.Equal(t, customer1.Id, int(customerResponse1["id"].(float64)))
	assert.Equal(t, customer1.Name, customerResponse1["name"])

	assert.Equal(t, customer2.Id, int(customerResponse2["id"].(float64)))
	assert.Equal(t, customer2.Name, customerResponse2["name"])

}

func TestUnauthorizedcustomer(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/customer", nil)

	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
