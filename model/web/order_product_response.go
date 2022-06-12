package web

import "time"

type OrderProductResponse struct {
	Id        int       `json:"id"`
	OrderId   int       `json:"orderId"`
	ProductId int       `json:"productId"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
