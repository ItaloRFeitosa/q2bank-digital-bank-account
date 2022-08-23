package core

import "context"

type ApplicationContext struct {
	AccountID uint
	UserID    uint
}

type CanBind[T any] interface {
	Bind(ApplicationContext) T
}

type UseCaseExecutor[In any, Out any] func(context.Context, In) (Out, error)
