package account

import "context"

type AccountDTO struct {
	ID       uint
	Balance  uint64
	UserID   uint
	Currency string
}

type NewAccountInput struct {
	Balance  uint64
	UserID   uint
	Currency string
}

type Repository interface {
	Insert(ctx context.Context, in NewAccountInput) (uint, error)
}

type UseCase interface {
	OpenAccount(ctx context.Context, userID uint) (uint, error)
}
