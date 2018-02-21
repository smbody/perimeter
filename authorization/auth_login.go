package authorization

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/smbody/perimeter/config"
)

func authLogin(rw http.ResponseWriter, req *http.Request) {

	// Устанавливаем набор параметров для токена
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    "test",
	}

	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен нашим секретным ключем
	tokenString, _ := token.SignedString(config.SigningSecretKey)

	// Отдаем токен клиенту
	rw.Write([]byte(tokenString))
}
