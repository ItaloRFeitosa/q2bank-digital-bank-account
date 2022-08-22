package transaction

import (
	"errors"

	"github.com/italorfeitosa/q2bank-digital-bank-account/common/error_builder"
)

var ErrSellerCannotCashOut = errors.New("SELLER_CANNOT_CASH_OUT")

var ErrBalanceNotEnough = errors.New("BALANCE_NOT_ENOUGH")

var ErrTransferUnauthorized = errors.New("UNAUTHORIZED")

var TransactionError = error_builder.OfBusiness.WithContext("Transaction")

var ErrSameAccount = TransactionError.AddReason("can't perform transfer between same accounts").Build()

var ErrAuthorizerWrongFormat = TransactionError.AddReason("authorizer returned wrong format").Build()

var ErrAuthorizerHasError = TransactionError.AddReason("authorizer returned with error").Build()
