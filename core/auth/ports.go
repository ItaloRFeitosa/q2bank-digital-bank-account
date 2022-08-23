package auth

import (
	"context"
	"time"

	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
)

type Service interface {
	SignUp(ctx context.Context, in user.RegisterUserInput) (SignOutput, error)
	SignIn(ctx context.Context, in SignInInput) (SignOutput, error)
	VerifyAndDecodeToken(token string) (SignData, error)
}

type SignOutput struct {
	Type        string    `json:"type"`
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
	IssuedAt    time.Time `json:"issued_at"`
}

type Repository interface {
	FindUserByEmail(ctx context.Context, email string) (UserData, error)
}

type UserData struct {
	ID        uint
	AccountID uint
	Password  string
}

type Token interface {
	Sign(data SignData) (SignOutput, error)
	Verify(token string) (SignData, error)
}

type SignData struct {
	UserID    uint `json:"user_id"`
	AccountID uint `json:"account_id"`
}
