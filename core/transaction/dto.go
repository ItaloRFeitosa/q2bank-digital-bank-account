package transaction

import (
	"time"

	"github.com/italorfeitosa/q2bank-digital-bank-account/core"
)

type TransferInput struct {
	PayerID uint
	PayeeID uint   `json:"payee_id" binding:"required"`
	Amount  uint64 `json:"amount" binding:"required,gte=1"`
}

func (TransferInput) Bind(app core.ApplicationContext) TransferInput {
	return TransferInput{
		PayerID: app.AccountID,
	}
}

type TransferOutput struct {
	ID        uint      `json:"id"`
	PayerID   uint      `json:"payer_id"`
	PayeeID   uint      `json:"payee_id"`
	Amount    uint64    `json:"amount"`
	Status    string    `json:"status"`
	Reason    string    `json:"reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func newTransferOutput(transfer *Transfer) TransferOutput {
	return TransferOutput{
		ID:        transfer.ID,
		PayerID:   transfer.Payer.ID,
		PayeeID:   transfer.Payee.ID,
		Amount:    transfer.Amount,
		Status:    transfer.Status.String(),
		Reason:    transfer.Reason(),
		CreatedAt: transfer.CreatedAt,
	}
}
