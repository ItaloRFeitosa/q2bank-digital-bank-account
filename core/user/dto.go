package user

import "github.com/italorfeitosa/q2bank-digital-bank-account/core"

type RegisterUserInput struct {
	Kind     Kind   `json:"kind" binding:"required,oneof=CUSTOMER SELLER"`
	Document string `json:"document" binding:"required,numeric,min=11,max=14"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (RegisterUserInput) Bind(app core.ApplicationContext) RegisterUserInput {
	return RegisterUserInput{}
}
