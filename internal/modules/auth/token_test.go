package auth

import (
	"crypto/rsa"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/italorfeitosa/q2bank-digital-bank-account/internal/shared/key"
	"github.com/stretchr/testify/suite"
)

func TestTokenImpl(t *testing.T) {
	key.Load()
	suite.Run(t, new(TokenTestSuite))
}

type TokenTestSuite struct {
	suite.Suite
	private *rsa.PrivateKey
	public  *rsa.PublicKey
}

func (s *TokenTestSuite) SetupTest() {
	var err error
	s.private, err = jwt.ParseRSAPrivateKeyFromPEM(key.Private())
	if err != nil {
		panic(err)
	}
	s.public, err = jwt.ParseRSAPublicKeyFromPEM(key.Public())
	if err != nil {
		panic(err)
	}
}

func (s *TokenTestSuite) TestSign_Success() {

}
