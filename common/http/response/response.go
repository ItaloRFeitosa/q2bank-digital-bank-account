package response

import (
	"net/http"

	h "github.com/italorfeitosa/q2bank-digital-bank-account/common/http"
)

type response struct {
	statusCode int
	body       any
}

func (r response) StatusCode() int {
	return r.statusCode
}

func (r response) Body() any {
	return r.body
}

func success(statusCode int) func(body any) h.Response {
	return func(data any) h.Response {
		return response{statusCode, h.M{
			"data": data,
		}}
	}
}

func fail(statusCode int) func(err error) h.Response {
	return func(err error) h.Response {
		return response{statusCode, h.M{
			"error": err,
		}}
	}
}

var OK = success(http.StatusOK)
var Created = success(http.StatusCreated)
var BadRequest = fail(http.StatusBadRequest)
var UnprocessableEntity = fail(http.StatusUnprocessableEntity)
var InternalServerError = fail(http.StatusInternalServerError)
