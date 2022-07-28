package error_builder

import "fmt"

type ErrorType string

const (
	ValidationError ErrorType = "VALIDATION_ERROR"
	BusinessError   ErrorType = "BUSINESS_ERROR"
)

type AppError struct {
	Type    ErrorType `json:"type"`
	Context string    `json:"context,omitempty"`
	Reasons []string  `json:"reasons"`
}

func (e AppError) Error() string {
	typ := fmt.Sprintf("error of type %s", e.Type)
	reasons := fmt.Sprintf("with reason %s", e.Reasons[0])
	return fmt.Sprint(typ, " ", reasons)
}

type ErrorBuilder struct {
	err AppError
}

var OfValidation = ofType(ValidationError)
var OfBusiness = ofType(BusinessError)

func ofType(t ErrorType) ErrorBuilder {
	return ErrorBuilder{err: AppError{
		Type: t,
	}}
}

func (b ErrorBuilder) WithContext(ctx string) ErrorBuilder {
	b.err.Context = ctx
	return b
}

func (b ErrorBuilder) AddReason(reason string) ErrorBuilder {
	b.err.Reasons = append(b.err.Reasons, reason)
	return b
}

func (b ErrorBuilder) Build() AppError {
	return b.err
}
