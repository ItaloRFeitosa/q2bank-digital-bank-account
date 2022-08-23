package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/adapters"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/auth"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core/user"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	controller := newController(auth.NewService(db))
	r.Group("/auth").
		POST("/sign-up", controller.SignUp).
		POST("/sign-in", controller.SignIn)
}

type controller struct {
	SignUp gin.HandlerFunc
	SignIn gin.HandlerFunc
}

var signInControler = adapters.NewGinAdapterBuilder[auth.SignInInput, auth.SignOutput]()
var signUpControler = adapters.NewGinAdapterBuilder[user.RegisterUserInput, auth.SignOutput]()

func newController(service auth.Service) *controller {
	return &controller{
		SignUp: signUpControler.
			Bind(user.RegisterUserInput{}).
			ThenExecute(service.SignUp).
			AndRespondWith(http.StatusOK),
		SignIn: signInControler.
			Bind(auth.SignInInput{}).
			ThenExecute(service.SignIn).
			AndRespondWith(http.StatusOK),
	}
}
