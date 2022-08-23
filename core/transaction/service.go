package transaction

import "context"

type service struct {
	repo Repository
	auth Authorizer
}

func (s *service) Transfer(ctx context.Context, in TransferInput) (TransferOutput, error) {
	payer, err := s.repo.FindAccount(ctx, in.PayerID)
	if err != nil {
		return TransferOutput{}, err
	}

	payee, err := s.repo.FindAccount(ctx, in.PayeeID)
	if err != nil {
		return TransferOutput{}, err
	}

	transfer, err := NewTransfer(payer, payee, in.Amount)

	if err != nil {
		return TransferOutput{}, err
	}

	if err := s.auth.Authorize(ctx, transfer); err != nil {
		return TransferOutput{}, err
	}

	transfer.Do()

	if err = s.repo.InsertTransfer(ctx, transfer); err != nil {
		return TransferOutput{}, err
	}

	return newTransferOutput(transfer), err
}
