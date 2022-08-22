package user

import (
	"context"
)

type Service struct {
	repo   Repository
	crypto Crypto
}

func (s *Service) RegisterUser(ctx context.Context, in RegisterUserInput) (uint, error) {
	exists, err := s.repo.Exists(ctx, in.Document, in.Email)

	if err != nil {
		return 0, err
	}

	if exists {
		return 0, ErrUserAlreadyExists
	}

	if err := s.crypto.Hash(&in.Password); err != nil {
		return 0, err
	}

	id, err := s.repo.Register(ctx, in)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Service) CheckPassword(hash, password string) bool {
	return s.crypto.Compare(hash, password)
}
