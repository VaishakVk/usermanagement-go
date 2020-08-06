package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword Func
func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

// CompareHash Func
func CompareHash(hash string, password string) error {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error
}
