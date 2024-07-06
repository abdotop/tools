package kryptonite

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKryptonite(t *testing.T) {
	// Initialize kryptonite with a secret key and SHA-256 hash function
	secretKey := "supersecretkey"
	k, err := New(secretKey, sha256.New)
	assert.NoError(t, err)
	// Define a password and a salt
	password := "password123"
	salt := []byte("somesalt123456") // Ensure salt is at least 8 bytes

	// Test GenerateHash
	hashedPassword, err := k.GenerateHash(password, salt)
	assert.NoError(t, err, "Generating hash should not produce an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")

	// Test CompareHashAndPassword with correct password
	err = k.CompareHashAndPassword(hashedPassword, password, salt)
	assert.NoError(t, err, "Comparison should succeed with correct password")

	// Test CompareHashAndPassword with incorrect password
	err = k.CompareHashAndPassword(hashedPassword, "wrongpassword", salt)
	assert.Error(t, err, "Comparison should fail with incorrect password")
	assert.Equal(t, "password does not match", err.Error(), "Error message should indicate mismatch")

	// Test GenerateHash with short salt
	shortSalt := []byte("short")
	_, err = k.GenerateHash(password, shortSalt)
	assert.Error(t, err, "Should error on short salt")
	assert.Equal(t, "salt too short, must be at least 8 bytes", err.Error(), "Error message should indicate short salt")
}
