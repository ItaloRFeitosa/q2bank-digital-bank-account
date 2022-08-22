package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/api/response"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/transaction"
)

type Handler struct {
	service transaction.UseCase
}

func (h *Handler) Transfer(c *gin.Context) {
	in := transaction.TransferInput{}
	if err := c.ShouldBindJSON(&in); err != nil {
		response.ValidationError(c, err)
		return
	}

	in.PayerID = c.MustGet("account_id").(uint)

	res, err := h.service.Transfer(c.Request.Context(), in)
	if err != nil {
		response.ServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Data(res))
}
