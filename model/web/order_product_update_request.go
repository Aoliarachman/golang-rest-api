package web

type OrderProductUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required" json:"orderId"`
	ProductId int `validate:"required" json:"productId"`
	Qty       int `validate:"required" json:"qty"`
	Price     int `validate:"required" json:"price"`
	Amount    int `validate:"required" json:"amount"`
}
