package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/api/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/transaction"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) *Handler {
	s := transaction.NewService(db)

	return &Handler{s}
}

func SetupRoutes(r *gin.Engine, a *auth.Handler, h *Handler) {

	r.POST("/transfer", a.EnsureAuthorized, h.Transfer)
}
