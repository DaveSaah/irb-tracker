package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// getHashPassword returns a hashed password using bcrypt and default cost
func GetHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Unable to hash password: %s", err)
		return "", err
	}

	return string(hash), nil
}
