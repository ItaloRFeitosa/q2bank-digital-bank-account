package transaction

import "context"

type Service struct {
	repo Repository
	auth Authorizer
}

func (s *Service) Transfer(ctx context.Context, in TransferInput) (TransferOutput, error) {
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

	return TransferOutput{
		ID:        transfer.ID,
		PayerID:   transfer.Payer.ID,
		PayeeID:   transfer.Payee.ID,
		Amount:    transfer.Amount,
		Status:    transfer.Status.String(),
		Reason:    transfer.Reason(),
		CreatedAt: transfer.CreatedAt,
	}, err
}
