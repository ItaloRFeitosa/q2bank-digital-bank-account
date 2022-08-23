package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/response"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
)

type Handler struct {
	service auth.Service
}

func (h *Handler) SignUp(c *gin.Context) {
	in := user.RegisterUserInput{}
	if err := c.ShouldBindJSON(&in); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.service.SignUp(c.Request.Context(), in)
	if err != nil {
		response.ServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.Data(res))
}

func (h *Handler) SignIn(c *gin.Context) {
	in := auth.SignInInput{}
	if err := c.ShouldBindJSON(&in); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.service.SignIn(c.Request.Context(), in)
	if err != nil {
		response.ServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Data(res))
}
