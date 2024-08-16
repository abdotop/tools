package kryptonite

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

// Kryptonite is a struct that contains the secret key and the hash function
type Kryptonite struct {
	secretKey     string
	hash_function func() hash.Hash
	errChan       chan error // Channel to send errors
}

// New creates a new Kryptonite instance with the secret key and the hash function ex: sha256
func New(secretKey string, h func() hash.Hash) (*Kryptonite, error) {
	if len(secretKey) < 8 { // Minimum length check for the secret key
		return nil, errors.New("secret key too short, must be at least 8 characters")
	}
	return &Kryptonite{
		secretKey:     secretKey,
		hash_function: h,
		errChan:       make(chan error),
	}, nil
}

// Hash Password With Secret now uses PBKDF2 for better security
func (k *Kryptonite) GenerateHash(password string, salt []byte) (string, error) {
	if len(salt) < 8 { // Ensuring that the salt is of adequate length
		return "", errors.New("salt too short, must be at least 8 bytes")
	}
	// Using PBKDF2 with HMAC-SHA-256, 10000 iterations, and 32-byte key length
	key := pbkdf2.Key([]byte(k.secretKey), salt, 10000, 32, k.hash_function)
	h := hmac.New(sha256.New, key)
	_, err := h.Write([]byte(password))
	if err != nil {
		k.errChan <- err
		return "", errors.New("error writing password: " + err.Error())
	}
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	return hashedPassword, nil
}

func (k *Kryptonite) CompareHashAndPassword(hashedPassword, password string, salt []byte) error {
	newHash, err := k.GenerateHash(password, salt)
	if err != nil {
		return errors.New("failed to hash password: " + err.Error())
	}
	if len(hashedPassword) != len(newHash) {
		err = errors.New("invalid hash length")
		return err
	}
	if subtle.ConstantTimeCompare([]byte(hashedPassword), []byte(newHash)) != 1 {
		err = errors.New("password does not match")
		return err
	}
	return nil
}

func (k *Kryptonite) OnError(callback func(error)) {
	go func() {
		for err := range k.errChan { // Correctly range over the channel
			if err != nil {
				callback(err) // Call the callback function with the error
			}
		}
	}()
}
