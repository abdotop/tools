package key

import (
	"math/rand"
	"time"
)

// Characters sets
const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	numbers          = "0123456789"
	specialChars     = "!@#$%^&*()-_+=[]{}|;:,.<>/?"
)

// GeneratePassword creates a random password with the specified length and complexity.
func Generate(length int, includeNumbers bool, includeSpecialChars bool) string {
	var passwordChars string
	passwordChars += uppercaseLetters + lowercaseLetters // Always include letters

	if includeNumbers {
		passwordChars += numbers
	}
	if includeSpecialChars {
		passwordChars += specialChars
	}

	// Create a new local random generator
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	password := make([]byte, length)
	for i := range password {
		password[i] = passwordChars[r.Intn(len(passwordChars))]
	}

	return string(password)
}
