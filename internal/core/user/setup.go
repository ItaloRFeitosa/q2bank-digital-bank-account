package user

import "gorm.io/gorm"

func NewService(db *gorm.DB) *Service {

	return &Service{repo: &Repo{db}, crypto: &CryptoImpl{}}
}
