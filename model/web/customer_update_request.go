package web

type CustomerUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,max=200,min=1" json:"name"`
	Address     string `validate:"required,max=100,min=1" json:"address"`
	Email       string `validate:"required,max=200,min=1" json:"email"`
	PhoneNumber string `validate:"required,max=20,min=1" json:"phonenumber"`
}
