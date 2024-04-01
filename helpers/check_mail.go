package helpers

import "regexp"

// IsMailValid checks if the email is valid or not.
func IsMailValid(email string) bool {
	return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email) || email == ""
}
