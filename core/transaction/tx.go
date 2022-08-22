package transaction

import (
	"context"

	"gorm.io/gorm"
)

type TxProxy struct {
	*Service
	db *gorm.DB
}

func (s *TxProxy) Transfer(ctx context.Context, in TransferInput) (TransferOutput, error) {
	tx := s.db.Begin()
	out, err := s.withTx(tx).Transfer(ctx, in)
	if err != nil {
		tx.Rollback()
		return out, err
	}
	tx.Commit()
	return out, err
}
