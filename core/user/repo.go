package user

import (
	"context"
	"errors"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/db"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Register(ctx context.Context, in RegisterUserInput) (uint, error) {
	user := &db.User{
		Kind:     in.Kind.String(),
		Document: in.Document,
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}

	err := r.db.WithContext(ctx).Create(&user).Error

	return user.ID, err
}

func (r *Repo) Exists(ctx context.Context, document string, email string) (bool, error) {
	user := &db.User{}
	err := r.db.WithContext(ctx).
		Where("document = ?", document).
		Or("email = ?", email).
		First(&user).Error

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}
