package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/validation"
	"github.com/italorfeitosa/q2bank-digital-bank-account/common/error_builder"
)

func Data(data any) gin.H {
	return gin.H{
		"data": data,
	}
}

func Error(err error) gin.H {
	return gin.H{
		"error": err,
	}
}

func ValidationError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error(validation.Error(errs)))
		return
	}

	InternalServerError(c, err)
}
func ServiceError(c *gin.Context, err error) {
	if err, ok := err.(error_builder.AppError); ok {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Error(err))
		return
	}

	InternalServerError(c, err)

}

func InternalServerError(c *gin.Context, err error) {
	log.Print(err.Error())
	c.AbortWithStatus(http.StatusInternalServerError)
}
