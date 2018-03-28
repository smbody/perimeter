package config

import (
	"crypto/rand"
	"fmt"
)

const (
	Port = 8400
)

const (
	AccessTokenLifeMinutes = 15
	TokenSize              = 16
	SaltSize               = 8
)

func GetSecretApp(appId string) string {
	return "gds64okey"
}

func GetPasswordHash(password string) string {
	return password
}

func CreateToken() string {
	return randomString(TokenSize)
}

func CreateSalt() string {
	return randomString(SaltSize)
}

func randomString(size int16) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
