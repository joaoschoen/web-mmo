package security

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func Decrypt(hash []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err
}
