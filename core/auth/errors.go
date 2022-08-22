package auth

import "github.com/italorfeitosa/q2bank-digital-bank-account/common/error_builder"

var AuthError = error_builder.
	OfBusiness.
	WithContext("Auth")

var ErrInvalidToken error = AuthError.
	AddReason("invalid token").
	Build()

var ErrUserNotFound error = AuthError.
	AddReason("user not found").
	Build()

var ErrSignInError error = AuthError.
	AddReason("can't sign in").
	Build()
