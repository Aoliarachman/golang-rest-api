package web

import "time"

type OrdersCreateRequest struct {
	Id          int
	OrderDate   time.Time
	CustomerId  int
	TotalAmount int
}
