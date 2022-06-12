package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required,max=200,min=1" json:"orderId"`
	Price      int    `validate:"required,max=100,min=1" json:"productId"`
	CategoryId int    `validate:"required,max=200,min=1" json:"qty"`
}
