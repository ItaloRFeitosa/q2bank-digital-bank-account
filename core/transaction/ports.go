package transaction

import (
	"context"
	"time"
)

type TransferInput struct {
	PayerID uint
	PayeeID uint   `json:"payee_id" binding:"required"`
	Amount  uint64 `json:"amount" binding:"required,gte=1"`
}

type TransferOutput struct {
	ID        uint      `json:"id"`
	PayerID   uint      `json:"payer_id"`
	PayeeID   uint      `json:"payee_id"`
	Amount    uint64    `json:"amount"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	InsertTransfer(context.Context, *Transfer) error
	FindAccount(ctx context.Context, id uint) (*Account, error)
}

type Authorizer interface {
	Authorize(context.Context, *Transfer) error
}

type UseCase interface {
	Transfer(context.Context, TransferInput) (TransferOutput, error)
}
