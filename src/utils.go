package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func comparePasswords(target, src string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(target), []byte(src))
	return err == nil
}
