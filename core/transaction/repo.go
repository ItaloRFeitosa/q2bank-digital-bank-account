package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/db"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/user"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func mapTransaction(t *Transaction) db.Transaction {
	return db.Transaction{
		Type:      t.Type.String(),
		Channel:   t.Channel.String(),
		Amount:    t.Amount,
		AccountID: t.AccountID,
	}
}

func (r *Repo) InsertTransfer(ctx context.Context, t *Transfer) error {
	sess := r.db.WithContext(ctx)
	transfer := db.Transfer{
		Status:  t.Status.String(),
		Amount:  t.Amount,
		PayeeID: t.Payee.ID,
		PayerID: t.Payer.ID,
	}

	if t.Done() {
		if err := sess.Model(&db.Account{}).Where("id = ?", t.Payee.ID).Update("balance", t.Payee.Balance).Error; err != nil {
			return err
		}
		if err := sess.Model(&db.Account{}).Where("id = ?", t.Payer.ID).Update("balance", t.Payer.Balance).Error; err != nil {
			return err
		}
		transfer.PayeeTx = mapTransaction(t.PayeeTx)
		transfer.PayerTx = mapTransaction(t.PayerTx)
	}

	if t.Failed() {
		transfer.Reason = t.Reason()
	}

	if err := sess.Create(&transfer).Error; err != nil {
		return err
	}

	t.ID = transfer.ID
	t.CreatedAt = transfer.CreatedAt.UTC()

	return nil
}

func (r *Repo) FindAccount(ctx context.Context, id uint) (*Account, error) {
	userFound := db.User{}
	err := r.db.WithContext(ctx).Joins("Account").First(&userFound, "Account.id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, TransactionError.AddReason(fmt.Sprintf("account with id: %d not found", id)).Build()
		}
		return nil, err
	}

	return &Account{
		Balance:  userFound.Account.Balance,
		ID:       userFound.Account.ID,
		UserKind: user.Kind(userFound.Kind),
	}, nil
}
