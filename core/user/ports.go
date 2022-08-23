package user

import (
	"context"
)

type Service interface {
	RegisterUser(ctx context.Context, in RegisterUserInput) (uint, error)
	CheckPassword(hash, password string) bool
}

type Repository interface {
	Exists(ctx context.Context, document string, email string) (bool, error)
	Register(ctx context.Context, in RegisterUserInput) (uint, error)
}

type Crypto interface {
	Hash(password *string) error
	Compare(hash, password string) bool
}
