package user

import "gorm.io/gorm"

func NewService(db *gorm.DB) *service {

	return &service{repo: &Repo{db}, crypto: &CryptoImpl{}}
}
