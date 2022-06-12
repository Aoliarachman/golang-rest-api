package domain

import "time"

type Orders struct {
	Id          int
	OrderDate   time.Time
	CustomerId  int
	TotalAmount int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
