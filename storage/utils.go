package storage

import (
	"golang.org/x/crypto/bcrypt"
)

// doesn't belong to storage package, place to some utils?
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
