package exception

import "fmt"

type ExceptionType string

const (
	ValidationException ExceptionType = "VALIDATION_EXCEPTION"
	BusinessException   ExceptionType = "BUSINESS_EXCEPTION"
)

type Exception struct {
	Type    ExceptionType `json:"type"`
	Context string        `json:"context,omitempty"`
	Reasons []string      `json:"reasons"`
}

func (e Exception) Error() string {
	typ := fmt.Sprintf("exception of type %s", e.Type)
	reasons := fmt.Sprintf("with reason %s", e.Reasons[0])
	return fmt.Sprint(typ, " ", reasons)
}

type Builder struct {
	err Exception
}

var OfValidation = ofType(ValidationException)
var OfBusiness = ofType(BusinessException)

func ofType(t ExceptionType) Builder {
	return Builder{err: Exception{
		Type: t,
	}}
}

func (b Builder) WithContext(ctx string) Builder {
	b.err.Context = ctx
	return b
}

func (b Builder) AddReason(reason string) Builder {
	b.err.Reasons = append(b.err.Reasons, reason)
	return b
}

func (b Builder) Build() Exception {
	return b.err
}
