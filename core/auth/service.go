package auth

import (
	"context"

	"github.com/italorfeitosa/q2bank-digital-bank-account/core/account"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
)

type service struct {
	users    user.Service
	accounts account.UseCase
	repo     Repository
	token    Token
}

func (s *service) SignUp(ctx context.Context, in user.RegisterUserInput) (SignOutput, error) {
	userID, err := s.users.RegisterUser(ctx, in)

	if err != nil {
		return SignOutput{}, err
	}

	accountID, err := s.accounts.OpenAccount(ctx, userID)

	if err != nil {
		return SignOutput{}, err
	}

	return s.token.Sign(SignData{userID, accountID})
}

func (s *service) SignIn(ctx context.Context, in SignInInput) (SignOutput, error) {
	u, err := s.repo.FindUserByEmail(ctx, in.Email)

	if err != nil {
		return SignOutput{}, err
	}

	if ok := s.users.CheckPassword(u.Password, in.Password); !ok {
		return SignOutput{}, ErrSignInError
	}

	return s.token.Sign(SignData{u.ID, u.AccountID})
}

func (s *service) VerifyAndDecodeToken(token string) (SignData, error) {
	return s.token.Verify(token)
}
