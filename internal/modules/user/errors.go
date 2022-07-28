package user

import "github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/error_builder"

var UserError = error_builder.
	OfBusiness.
	WithContext("User")

var ErrUserAlreadyExists error = UserError.
	AddReason("can't register user with given email or document").
	Build()
