package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Middleware) EnsureAuthorization(c *gin.Context) {
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
