package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/q2bank-digital-bank-account/adapters/api/response"
	"github.com/italorfeitosa/q2bank-digital-bank-account/common/exception"
	"github.com/italorfeitosa/q2bank-digital-bank-account/core"
)

type GinAdapter[In any, Out any] struct {
	input             core.CanBind[In]
	executeUseCase    core.UseCaseExecutor[In, Out]
	defaultStatusCode int
}

type GinAdapterBuilder[In any, Out any] struct {
	adapter *GinAdapter[In, Out]
}

func NewGinAdapterBuilder[In any, Out any]() *GinAdapterBuilder[In, Out] {
	return &GinAdapterBuilder[In, Out]{&GinAdapter[In, Out]{}}
}

func (builder *GinAdapterBuilder[In, Out]) Bind(input core.CanBind[In]) *GinAdapterBuilder[In, Out] {
	builder.adapter.input = input
	return builder
}

func (builder *GinAdapterBuilder[In, Out]) ThenExecute(executeUseCase core.UseCaseExecutor[In, Out]) *GinAdapterBuilder[In, Out] {
	builder.adapter.executeUseCase = executeUseCase
	return builder
}

func (builder *GinAdapterBuilder[In, Out]) AndRespondWith(defaultStatusCode int) gin.HandlerFunc {
	builder.adapter.defaultStatusCode = defaultStatusCode
	return builder.adapter.Handle
}

func (g *GinAdapter[In, Out]) Handle(c *gin.Context) {
	in := g.input.Bind(core.ApplicationContext{
		AccountID: c.GetUint("account_id"),
		UserID:    c.GetUint("user_id"),
	})

	if err := c.ShouldBindJSON(&in); err != nil {
		response.ValidationError(c, err)
		return
	}

	out, err := g.executeUseCase(c.Request.Context(), in)

	switch err.(type) {
	case nil:
		c.JSON(g.defaultStatusCode, response.Data(out))
		return
	case exception.Exception:
		response.ServiceError(c, err)
		return
	default:
		response.InternalServerError(c, err)
		return
	}
}
