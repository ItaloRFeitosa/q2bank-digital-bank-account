package user

import "golang.org/x/crypto/bcrypt"

type CryptoImpl struct{}

func (c *CryptoImpl) Hash(password *string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*password), 10)

	if err != nil {
		return err
	}

	*password = string(hashed)

	return nil
}

func (c *CryptoImpl) Compare(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
