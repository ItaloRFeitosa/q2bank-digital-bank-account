package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/modules/auth"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) *Handler {
	s := auth.NewService(db)

	return &Handler{s}
}

func SetupRoutes(r *gin.Engine, h *Handler) {

	r.Group("/auth").
		POST("/sign-up", h.SignUp).
		POST("/sign-in", h.SignIn)
}
