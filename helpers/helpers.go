// Package helpers defines the utility functions that are used across the application.
package helpers

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// IsMailValid checks if the email is valid or not.
func IsMailValid(email string) bool {
	return regexp.MustCompile(
		`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`,
	).MatchString(email) || email == ""
}

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
