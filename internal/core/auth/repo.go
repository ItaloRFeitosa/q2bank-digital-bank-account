package auth

import (
	"context"
	"errors"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/db"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (r *Repo) FindUserByEmail(ctx context.Context, email string) (UserData, error) {
	user := &db.User{
		Email: email,
	}

	result := r.db.WithContext(ctx).Where(&user).Joins("Account").First(&user)

	if result.Error == nil {
		return UserData{
			ID:        user.ID,
			Password:  user.Password,
			AccountID: user.Account.ID,
		}, nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return UserData{}, ErrUserNotFound
	}

	return UserData{}, result.Error
}
