package db

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Kind     string `gorm:"not null;size:8"`
	Document string `gorm:"not null;size:14;index:idx_user_document,unique"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;size:254;index:idx_user_email,unique"`
	Password string `gorm:"not null"`
	Account  Account
}

type Account struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Balance  uint64 `gorm:"not null"`
	Currency string `gorm:"not null"`
}

type Transaction struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null;size:8"`
	Channel   string `gorm:"not null;size:8"`
	Amount    uint64 `gorm:"not null"`
	AccountID uint   `gorm:"not null"`
	Account   Account
	CreatedAt time.Time
}

type Transfer struct {
	ID        uint   `gorm:"primaryKey"`
	Status    string `gorm:"not null;size:8"`
	Reason    string
	Amount    uint64 `gorm:"not null"`
	PayeeID   uint   `gorm:"not null"`
	Payee     Account
	PayeeTxID sql.NullInt64
	PayeeTx   Transaction
	PayerID   uint `gorm:"not null"`
	Payer     Account
	PayerTxID sql.NullInt64
	PayerTx   Transaction
	CreatedAt time.Time
}
