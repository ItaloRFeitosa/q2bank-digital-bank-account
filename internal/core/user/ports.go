package user

import "context"

type RegisterUserInput struct {
	Kind     Kind   `json:"kind" binding:"required,oneof=CUSTOMER SELLER"`
	Document string `json:"document" binding:"required,numeric,min=11,max=14"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UseCase interface {
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
