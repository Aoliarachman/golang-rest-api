package web

type OrdersUpdateRequest struct {
	Id        int `validate:"required"`
	OrderId   int `validate:"required,max=200,min=1" json:"orderId"`
	ProductId int `validate:"required,max=100,min=1" json:"productId"`
	Qty       int `validate:"required,max=200,min=1" json:"qty"`
	Price     int `validate:"required,max=20,min=1" json:"price"`
	Amount    int `validate:"required,max=20,min=1" json:"amount"`
}
