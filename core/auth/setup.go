package auth

import (
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/account"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	return &TxProxy{db: db, service: &service{
		user.NewService(db),
		account.NewService(db),
		&Repo{db},
		NewToken(),
	}}
}

func (s *TxProxy) withTx(tx *gorm.DB) Service {
	return &service{
		user.NewService(tx),
		account.NewService(tx),
		&Repo{tx},
		s.token,
	}
}
