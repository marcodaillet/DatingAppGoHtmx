package main

import (
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