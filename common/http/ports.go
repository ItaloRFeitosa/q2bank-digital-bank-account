package http

import "context"

type Request interface {
	Body() any
}

type Response interface {
	StatusCode() int
	Body() any
}

type M = map[string]any

type ControllerFunc = func(ctx context.Context, req Request) (Response, error)
