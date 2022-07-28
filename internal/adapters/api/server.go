package api

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/api/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/api/transaction"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/api/validation"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/adapters/db"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/config"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/key"
	"gorm.io/gorm"
)

func InitServer() {
	key.Load()
	config.Load()
	conn := db.DB()
	r := gin.Default()
	validation.SetupValidator()
	SetupRoutes(r, conn)

	r.Run(":3333")

}

func SetupRoutes(r *gin.Engine, conn *gorm.DB) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it's running",
		})
	})
	a := auth.NewHandler(conn)
	t := transaction.NewHandler(conn)
	auth.SetupRoutes(r, a)
	transaction.SetupRoutes(r, a, t)
}
