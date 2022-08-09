package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(password string, inputPassword []byte) bool {
	byteHash := []byte(password)

	if err := bcrypt.CompareHashAndPassword(byteHash, inputPassword); err != nil {
		return false
	}

	return true
}
