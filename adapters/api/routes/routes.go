package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/middleware"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/routes/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/routes/transaction"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/db"
)

func Setup(r *gin.Engine) {
	conn := db.DB()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "it's running",
		})
	})
	auth.Setup(r, conn)
	m := middleware.New(conn)
	transaction.Setup(r, conn, m)
}
