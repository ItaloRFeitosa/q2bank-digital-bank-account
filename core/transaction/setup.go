package transaction

import "gorm.io/gorm"

func NewService(db *gorm.DB) Service {
	return &transactionProxy{db: db, service: &service{
		NewRepository(db),
		NewAuthorizer(),
	}}
}

func (s *transactionProxy) withTx(tx *gorm.DB) Service {
	return &service{
		NewRepository(tx),
		s.auth,
	}
}
