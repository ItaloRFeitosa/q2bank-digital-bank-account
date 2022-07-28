package account

import (
	"context"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/db"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Insert(ctx context.Context, in NewAccountInput) (uint, error) {
	account := &db.Account{
		UserID:   in.UserID,
		Balance:  in.Balance,
		Currency: in.Currency,
	}
	result := r.db.WithContext(ctx).Create(&account)

	return account.ID, result.Error
}
