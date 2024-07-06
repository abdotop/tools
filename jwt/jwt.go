package jwt

import (
	"crypto/rsa"
	"encoding/base64"
	"errors"
	"os"

	"time"

	"github.com/abdotop/tools/jwt/utils"
	"github.com/golang-jwt/jwt/v4"
)

// jwt_tools est une structure qui contient la clé privée et la clé publique.
type jwt_tools struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	errChan    chan error // Channel to send errors
	exp        int64
}

// New crée une nouvelle instance de jwt_tools.
func New(exp uint8) *jwt_tools {
	return &jwt_tools{
		errChan: make(chan error),
		exp:     time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	}
}

// LoadPrivateKeyFromEnv charge la clé privée depuis l'environnement.
func (j *jwt_tools) LoadPrivateKeyFromEnv(key string) error {
	if key == "" {
		err := errors.New("key is empty")
		j.errChan <- err
		return err
	}
	base64Key := os.Getenv(key)
	if base64Key == "" {
		err := errors.New("key not found")
		j.errChan <- err
		return err
	}
	keyData, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		j.errChan <- err
		return err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		j.errChan <- err
		return err
	}
	j.privateKey = privateKey
	return nil
}

// LoadPublicKeyFromEnv charge la clé publique depuis l'environnement.
func (j *jwt_tools) LoadPublicKeyFromEnv(key string) error {
	if key == "" {
		err := errors.New("key is empty")
		j.errChan <- err
		return err
	}
	base64Key := os.Getenv(key)
	if base64Key == "" {
		err := errors.New("key not found")
		j.errChan <- err
		return err
	}
	keyData, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		j.errChan <- err
		return err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		j.errChan <- err
		return err
	}
	j.publicKey = publicKey
	return nil
}

func (j *jwt_tools) LoadPrivateKeyFromSecretsManager(secretName string) error {
	secret, err := utils.GetSecret(secretName)
	if err != nil {
		j.errChan <- err
		return err
	}
	keyData, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		j.errChan <- err
		return err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		j.errChan <- err
		return err
	}
	j.privateKey = privateKey
	return nil
}

func (j *jwt_tools) LoadPublicKeyFromSecretsManager(secretName string) error {
	secret, err := utils.GetSecret(secretName)
	if err != nil {
		j.errChan <- err
		return err
	}
	keyData, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		j.errChan <- err
		return err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		j.errChan <- err
		return err
	}
	j.publicKey = publicKey
	return nil
}

// GenerateToken génère un nouveau token JWT.
func (j *jwt_tools) GenerateToken(data interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"data": data,
		"exp":  j.exp,
	})

	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		j.errChan <- err
		return "", err
	}
	return tokenString, nil
}

// ValidateToken valide un token JWT et retourne les claims s'il est valide.
func (j *jwt_tools) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			err := jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			j.errChan <- err
			return nil, err
		}
		return j.publicKey, nil
	})

	if err != nil {
		j.errChan <- err
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err := jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
		j.errChan <- err
		return nil, err
	}
}

func (j *jwt_tools) OnError(callback func(error)) {
	go func() {
		for err := range j.errChan { // Correctly range over the channel
			if err != nil {
				callback(err) // Call the callback function with the error
			}
		}
	}()
}
