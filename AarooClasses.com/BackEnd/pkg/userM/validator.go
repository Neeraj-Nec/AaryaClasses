package userM

import (
	"net/mail"
	"regexp"
)

func validateUserName(username string) bool {
	if username == "" {
		return false
	} 	
	return true
} 
func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func validatePassword(password string) bool {
	// Regular expression to match the password criteria
	hasSpecialChar := regexp.MustCompile(`[[:punct:]]`).MatchString(password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	isMinLength := len(password) >= 8
	return hasSpecialChar && hasUppercase && isMinLength && hasLowerCase
}
