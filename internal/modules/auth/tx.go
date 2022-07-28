package auth

import (
	"context"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/user"
	"gorm.io/gorm"
)

type TxProxy struct {
	*Service
	db *gorm.DB
}

func (s *TxProxy) SignUp(ctx context.Context, in user.RegisterUserInput) (*SignOutput, error) {
	tx := s.db.Begin()
	out, err := s.withTx(tx).SignUp(ctx, in)
	if err != nil {
		tx.Rollback()
		return out, err
	}
	tx.Commit()
	return out, err
}
