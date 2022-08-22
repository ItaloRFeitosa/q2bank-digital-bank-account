package auth

import (
	"crypto/rsa"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/key"
)

const defaultTTL time.Duration = 15 * time.Minute

type TokenImpl struct {
	private *rsa.PrivateKey
	public  *rsa.PublicKey
	ttl     time.Duration
}

func NewToken() *TokenImpl {
	private, err := jwt.ParseRSAPrivateKeyFromPEM(key.Private())
	if err != nil {
		log.Fatal(err)
	}
	public, err := jwt.ParseRSAPublicKeyFromPEM(key.Public())
	if err != nil {
		log.Fatal(err)
	}
	return &TokenImpl{private, public, defaultTTL}
}

type Claims struct {
	jwt.RegisteredClaims
	SignData
}

func (t *TokenImpl) Sign(data SignData) (*SignOutput, error) {
	iat := time.Now().UTC()
	exp := iat.Add(t.ttl)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &Claims{
		SignData: data,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		}})

	signed, err := token.SignedString(t.private)
	if err != nil {
		return nil, err
	}

	return &SignOutput{
		Type:        "Bearer",
		AccessToken: signed,
		ExpiresAt:   exp,
		IssuedAt:    iat,
	}, nil
}

func (t *TokenImpl) Verify(tk string) (SignData, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return t.public, nil
	})

	if err != nil {
		return SignData{}, err
	}

	if !token.Valid {
		return SignData{}, ErrInvalidToken
	}

	return claims.SignData, nil
}
