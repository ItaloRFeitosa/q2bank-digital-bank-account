package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/response"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
)

type Handler struct {
	service auth.UseCase
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

func (h *Handler) EnsureAuthorized(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(header, " ")

	if len(token) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if token[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	data, err := h.service.VerifyAndDecodeToken(token[1])

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("account_id", data.AccountID)
	c.Set("user_id", data.UserID)

	c.Next()
}
