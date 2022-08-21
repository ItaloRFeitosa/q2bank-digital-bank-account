package account

import (
	"context"
)

type Service struct {
	repo Repository
}

func (s *Service) OpenAccount(ctx context.Context, userID uint) (uint, error) {
	return s.repo.Insert(ctx, OpenAccount(userID))
}
