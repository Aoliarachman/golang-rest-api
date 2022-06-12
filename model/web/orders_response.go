package web

import "time"

type OrdersResponse struct {
	Id          int       `json:"id"`
	OrderDate   time.Time `json:"orderDateId"`
	CustomerId  int       `json:"customerId"`
	TotalAmount int       `json:"totalAmount"`
}
