package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/db"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
func mapTransactionToDB(t *Transaction) db.Transaction {
	return db.Transaction{
		Type:      t.Type.String(),
		Channel:   t.Channel.String(),
		Amount:    t.Amount,
		AccountID: t.AccountID,
	}
}

func (r *repository) InsertTransfer(ctx context.Context, transfer *Transfer) error {
	sess := r.db.WithContext(ctx)
	transferToPersist := db.Transfer{
		Status:  transfer.Status.String(),
		Amount:  transfer.Amount,
		PayeeID: transfer.Payee.ID,
		PayerID: transfer.Payer.ID,
	}

	if transfer.Done() {
		if err := sess.Model(&db.Account{}).Where("id = ?", transfer.Payee.ID).Update("balance", transfer.Payee.Balance).Error; err != nil {
			return err
		}
		if err := sess.Model(&db.Account{}).Where("id = ?", transfer.Payer.ID).Update("balance", transfer.Payer.Balance).Error; err != nil {
			return err
		}
		transferToPersist.PayeeTx = mapTransactionToDB(transfer.PayeeTx)
		transferToPersist.PayerTx = mapTransactionToDB(transfer.PayerTx)
	}

	if transfer.Failed() {
		transferToPersist.Reason = transfer.Reason()
	}

	if err := sess.Create(&transferToPersist).Error; err != nil {
		return err
	}

	transfer.ID = transferToPersist.ID
	transfer.CreatedAt = transferToPersist.CreatedAt.UTC()

	return nil
}

func (r *repository) FindAccount(ctx context.Context, id uint) (*Account, error) {
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
