package validation

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/error_builder"
)

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func Error(errs validator.ValidationErrors) error {
	builder := error_builder.
		OfValidation.
		WithContext(strings.Split(errs[0].Namespace(), ".")[0])

	for _, err := range errs {
		builder = builder.AddReason(err.Field() + " must be " + err.Tag() + " " + err.Param())
	}
	return builder.Build()
}
