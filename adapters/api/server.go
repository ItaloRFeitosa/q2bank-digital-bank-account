package api

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/routes"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/validation"
	"github.com/italorfeitosa/q2bank-digital-bank-account/common/config"
	"github.com/italorfeitosa/q2bank-digital-bank-account/common/key"
)

func Setup() {
	key.Load()
	config.Load()
	r := gin.Default()
	validation.SetupValidator()
	routes.Setup(r)
	r.Run(":3333")
}
