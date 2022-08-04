package postgres

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("faild to encript passowrd: %w", err)
	}

	return string(hashed), nil
}

func isPasswordValid(pass, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err == nil
}
