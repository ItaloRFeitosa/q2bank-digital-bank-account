package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/adapters"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/middleware"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/transaction"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB, m *middleware.Middleware) {
	controller := newController(transaction.NewService(db))
	r.POST("/transfer", m.EnsureAuthorization, controller.Transfer)
}

type controller struct {
	Transfer gin.HandlerFunc
}

var transferController = adapters.NewGinAdapterBuilder[transaction.TransferInput, transaction.TransferOutput]()

func newController(service transaction.Service) *controller {
	return &controller{
		Transfer: transferController.
			Bind(transaction.TransferInput{}).
			ThenExecute(service.Transfer).
			AndRespondWith(http.StatusCreated),
	}
}
