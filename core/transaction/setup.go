package transaction

import "gorm.io/gorm"

func NewService(db *gorm.DB) *TxProxy {
	return &TxProxy{db: db, Service: &Service{
		&Repo{db},
		NewAuthorizer(),
	}}
}

func (s *TxProxy) withTx(tx *gorm.DB) *Service {
	return &Service{
		&Repo{tx},
		s.auth,
	}
}
