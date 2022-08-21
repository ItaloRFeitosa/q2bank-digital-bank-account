package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCrypoImpl(t *testing.T) {
	t.Run("given a password should hash password and be comparable", func(t *testing.T) {
		r := require.New(t)
		c := new(CryptoImpl)

		password := "12345678"
		toBeHashed := password

		err := c.Hash(&toBeHashed)
		r.Nil(err, "c.Hash should return nil")
		r.NotEqual(toBeHashed, password, "hashed value should be different from original value")

		ok := c.Compare(toBeHashed, password)

		r.True(ok, "password should match with hashed password")
	})

}
