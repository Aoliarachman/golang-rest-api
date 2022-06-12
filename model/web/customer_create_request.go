package web

type CustomerCreateRequest struct {
	Id          int
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Address     string `validate:"required,min=1,max=100" json:"address"`
	Email       string `validate:"required,min=1,max=200" json:"email"`
	PhoneNumber string `validate:"required,min=1,max=20" json:"phonenumber"`
}
