package auth

import "github.com/italorfeitosa/q2bank-digital-bank-account/core"

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (SignInInput) Bind(app core.ApplicationContext) SignInInput {
	return SignInInput{}
}
