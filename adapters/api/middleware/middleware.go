package middleware

import (
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/auth"
	"gorm.io/gorm"
)

type Middleware struct {
	service auth.Service
}

func New(db *gorm.DB) *Middleware {
	return &Middleware{auth.NewService(db)}
}
