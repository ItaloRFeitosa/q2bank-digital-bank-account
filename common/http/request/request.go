package request

import "github.com/italorfeitosa/q2bank-digital-bank-account/common/http"

type request struct {
	body any
}

func (r request) Body() any {
	return r.body
}

func New(b any) http.Request {
	return request{b}
}
