package config

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordSize = 72 // ограничение bcrypt
)

func GetHash(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}

func GetPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func PasswordValidate(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
