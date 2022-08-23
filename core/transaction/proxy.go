package transaction

import (
	"context"

	"gorm.io/gorm"
)

type transactionProxy struct {
	*service
	db *gorm.DB
}

func (s *transactionProxy) Transfer(ctx context.Context, in TransferInput) (TransferOutput, error) {
	tx := s.db.Begin()
	out, err := s.withTx(tx).Transfer(ctx, in)
	if err != nil {
		tx.Rollback()
		return out, err
	}
	tx.Commit()
	return out, err
}
