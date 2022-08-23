package request

import "github.com/italorfeitosa/q2bank-digital-bank-account/common/exception"

var ErrMalformedBody = exception.OfValidation.AddReason("malformed request").Build()
