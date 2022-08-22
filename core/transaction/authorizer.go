package transaction

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/italorfeitosa/q2bank-digital-bank-account/common/config"
)

type AuthorizerImpl struct {
	client *resty.Client
}

type AuthorizerResponse struct {
	Authorization bool `json:"authorization"`
}

func NewAuthorizer() *AuthorizerImpl {
	client := resty.New()
	client.SetBaseURL(config.AuthorizerUrl)
	return &AuthorizerImpl{client}
}

func (a *AuthorizerImpl) Authorize(ctx context.Context, t *Transfer) error {
	resp, err := a.client.R().
		SetContext(ctx).
		SetResult(AuthorizerResponse{}).
		Get("/")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return ErrAuthorizerHasError
	}

	data, ok := resp.Result().(*AuthorizerResponse)

	if !ok {
		return ErrAuthorizerWrongFormat
	}

	if data.Authorization {
		t.Authorize()
	}

	return nil
}
