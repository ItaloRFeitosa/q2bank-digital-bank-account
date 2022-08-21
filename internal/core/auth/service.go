package auth

import (
	"context"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/account"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/user"
)

type Service struct {
	users    user.UseCase
	accounts account.UseCase
	repo     Repository
	token    Token
}

func (s *Service) SignUp(ctx context.Context, in user.RegisterUserInput) (*SignOutput, error) {
	userID, err := s.users.RegisterUser(ctx, in)

	if err != nil {
		return nil, err
	}

	accountID, err := s.accounts.OpenAccount(ctx, userID)

	if err != nil {
		return nil, err
	}

	return s.token.Sign(SignData{userID, accountID})
}

func (s *Service) SignIn(ctx context.Context, in SignInInput) (*SignOutput, error) {
	u, err := s.repo.FindUserByEmail(ctx, in.Email)

	if err != nil {
		return nil, err
	}

	if ok := s.users.CheckPassword(u.Password, in.Password); !ok {
		return nil, ErrSignInError
	}

	return s.token.Sign(SignData{u.ID, u.AccountID})
}

func (s *Service) VerifyAndDecodeToken(token string) (SignData, error) {
	return s.token.Verify(token)
}
