package auth

import (
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/account"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/user"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) *TxProxy {
	return &TxProxy{db: db, Service: &Service{
		user.NewService(db),
		account.NewService(db),
		&Repo{db},
		NewToken(),
	}}
}

func (s *TxProxy) withTx(tx *gorm.DB) *Service {
	return &Service{
		user.NewService(tx),
		account.NewService(tx),
		&Repo{tx},
		s.token,
	}
}
