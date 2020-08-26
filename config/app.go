package config

import (
	"crypto/rand"
	"fmt"
)

const (
	Port = 8400
)

const (
	AccessTokenLifeMinutes  = 15
	RefreshTokenLifeMinutes = 15 * 100 // в сто раз больше - 15 минут против суток
	TokenSize               = 16
	SaltSize                = 8
)

var applications = map[string]string{
	"appLimsMobile": "Lims-gds64Okey", // md5 = 19888fab31544bf6bcdf1210936c68c1
	"appTest":       "gds64okey",      // md5 = 55b948f13d282a4e2d86adbc73e825f2
	"fastride":      "frq1447pdu",     // md5 = 439199b179d0a8258642981b288a6d8b
}

var apps map[string]string

func init() {
	apps = make(map[string]string)
	for app, key := range applications {
		apps[GetHash(app)] = key
	}
}

func randomString(size int16) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GetSecretApp(appId string) string {
	return apps[appId]
}

func CreateToken() string {
	return randomString(TokenSize)
}

func CreateSalt() string {
	return randomString(SaltSize)
}
