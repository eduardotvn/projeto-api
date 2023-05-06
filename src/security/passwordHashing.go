package security

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encodedPassword := base64.StdEncoding.EncodeToString(hashedPassword)
	return encodedPassword, nil
}

func ValidatePassword(password, hashedPassword string) (bool, error) {
	hashBytes, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
