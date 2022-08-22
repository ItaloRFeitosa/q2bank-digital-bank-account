package transaction

import (
	"time"

	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
)

type Type string

const (
	TypeCashOut Type = "CASH_OUT"
	TypeCashIn  Type = "CASH_IN"
)

func (t Type) String() string {
	return string(t)
}

type Channel string

const (
	ChannelTransfer Channel = "TRANSFER"
	ChannelDeposit  Channel = "DEPOSIT"
	ChannelWithdraw Channel = "WITHDRAW"
)

func (c Channel) String() string {
	return string(c)
}

type Transaction struct {
	Type      Type
	Channel   Channel
	AccountID uint
	Amount    uint64
}

type Account struct {
	ID       uint
	UserKind user.Kind
	Balance  uint64
}

func (a *Account) CashIn(amount uint64) {
	a.Balance += amount
}

func (a *Account) CashOut(amount uint64) error {
	if a.UserKind.IsSeller() {
		return ErrSellerCannotCashOut
	}

	if a.Balance < amount {
		return ErrBalanceNotEnough
	}

	a.Balance -= amount
	return nil
}

type TransferStatus string

const (
	TransferStatusPending TransferStatus = "PENDING"
	TransferStatusFailed  TransferStatus = "FAILED"
	TransferStatusDone    TransferStatus = "DONE"
)

func (t TransferStatus) String() string {
	return string(t)
}

type Transfer struct {
	ID         uint
	Amount     uint64
	Status     TransferStatus
	reason     error
	PayerTx    *Transaction
	Payee      *Account
	PayeeTx    *Transaction
	Payer      *Account
	Authorized bool
	CreatedAt  time.Time
}

func NewTransfer(payer *Account, payee *Account, amount uint64) (*Transfer, error) {
	if payer.ID == payee.ID {
		return nil, ErrSameAccount
	}

	return &Transfer{
		Amount:     amount,
		Status:     TransferStatusPending,
		Payer:      payer,
		PayerTx:    nil,
		Payee:      payee,
		PayeeTx:    nil,
		Authorized: false,
	}, nil
}

func (t *Transfer) Authorize() {
	t.Authorized = true
}

func (t *Transfer) Do() {
	if !t.Authorized {
		t.Fail(ErrTransferUnauthorized)
		return
	}

	err := t.Payer.CashOut(t.Amount)
	if err != nil {
		t.Fail(err)
		return
	}

	t.PayerTx = &Transaction{
		Type:      TypeCashOut,
		Channel:   ChannelTransfer,
		AccountID: t.Payer.ID,
		Amount:    t.Amount,
	}

	t.Payee.CashIn(t.Amount)
	t.PayeeTx = &Transaction{
		Type:      TypeCashIn,
		Channel:   ChannelTransfer,
		AccountID: t.Payee.ID,
		Amount:    t.Amount,
	}

	t.Status = TransferStatusDone
}

func (t *Transfer) Done() bool {
	return t.Status == TransferStatusDone
}

func (t *Transfer) Fail(err error) {
	t.PayeeTx = nil
	t.PayerTx = nil
	t.Status = TransferStatusFailed
	t.reason = err
}

func (t *Transfer) Failed() bool {
	return t.Status == TransferStatusFailed
}

func (t *Transfer) Reason() string {
	if t.reason != nil {
		return t.reason.Error()
	}

	return ""
}
