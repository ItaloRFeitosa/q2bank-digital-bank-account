package transaction

import (
	"context"
)

type Repository interface {
	InsertTransfer(context.Context, *Transfer) error
	FindAccount(ctx context.Context, id uint) (*Account, error)
}

type Authorizer interface {
	Authorize(context.Context, *Transfer) error
}

type Service interface {
	Transfer(context.Context, TransferInput) (TransferOutput, error)
}
