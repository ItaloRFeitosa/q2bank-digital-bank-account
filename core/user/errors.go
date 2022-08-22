package user

import "github.com/italorfeitosa/q2bank-digital-bank-account/common/exception"

var UserError = exception.
	OfBusiness.
	WithContext("User")

var ErrUserAlreadyExists error = UserError.
	AddReason("can't register user with given email or document").
	Build()
