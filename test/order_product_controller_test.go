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

func setupTestDBorder_product() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_rest_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouterorder_product(db *sql.DB) http.Handler {
	validate := validator.New()
	order_productRepository := repository.NewOrderProductRepository()
	order_productService := service.NewOrderProductService(order_productRepository, db, validate)
	order_productController := controller.NewOrderProductController(order_productService)
	router := app.NewRouter(order_productController)

	return middleware.NewAuthMiddleware(router)
}

func truncateOrderProduct(db *sql.DB) {
	db.Exec("TRUNCATE order_product")
}

func TestCreateOrderProduct(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/order_product", requestBody)
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

func TestCreateOrderProductFailed(t *testing.T) {

	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/order_product", requestBody)
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

func TestUpdateOrderProductSucces(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	order_productRespository := repository.NewOrderProductRepository()
	order_product := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		Id: 1515,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"Id" : "1515"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/order_product"+strconv.Itoa(order_product.Id), requestBody)
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
	assert.Equal(t, order_product.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestUpdateOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	order_productRespository := repository.NewOrderProductRepository()
	order_product := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		Id: 1515,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"id" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/order_product"+strconv.Itoa(order_product.Id), requestBody)
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

func TestGetOrderProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	order_productRespository := repository.NewOrderProductRepository()
	order_product := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		Id: 1515,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/order_product/"+strconv.Itoa(order_product.Id), nil)
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
	assert.Equal(t, order_product.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, order_product.OrderId, responseBody["data"].(map[string]interface{})["orderproduct"])

}

func TestGetOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/order_product/404", nil)
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

func TestDeleteOrderProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	order_productRespository := repository.NewOrderProductRepository()
	order_product := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		Id: 1515,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/order_product/"+strconv.Itoa(order_product.Id), nil)
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

func TestDeleteOrderProductFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categpries/404", nil)
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

func TestListOrderProductSuccess(t *testing.T) {
	db := setupTestDB()
	truncateOrderProduct(db)

	tx, _ := db.Begin()
	order_productRespository := repository.NewOrderProductRepository()
	order_product1 := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		Id: 1515,
	})
	order_product2 := order_productRespository.Save(context.Background(), tx, domain.OrderProduct{
		OrderId: 20202,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/order_product", nil)
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

	var order_product = responseBody["data"].([]interface{})

	order_productResponse1 := order_product[0].(map[string]interface{})
	order_productResponse2 := order_product[1].(map[string]interface{})

	assert.Equal(t, order_product1.Id, int(order_productResponse1["id"].(float64)))
	assert.Equal(t, order_product1.OrderId, order_productResponse1["OrderId"])

	assert.Equal(t, order_product2.Id, int(order_productResponse2["id"].(float64)))
	assert.Equal(t, order_product2.OrderId, order_productResponse2["OrderId"])

}

func TestUnauthorizedOrderProduct(t *testing.T) {
	db := setupTestDB()
	truncateCustomer(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/order_product", nil)

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
