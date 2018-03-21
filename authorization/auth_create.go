package authorization

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/smbody/perimeter/config"
)

func authCreate(rw http.ResponseWriter, req *http.Request) {

	/// для имитации выдачи токена
	/// !!!

	request := new(AuthRequest)

	// для начала определим app
	appId := req.URL.Query().Get("appId")
	if appId != "" {
		request.AppId = appId
	} else {
		http.Error(rw, config.HttpErrorBadRequestAppId, http.StatusBadRequest)
		return
	}
	if request.Secret() == "" {
		http.Error(rw, config.HttpErrorBadRequestApp, http.StatusBadRequest)
		return
	}
	fmt.Println("create token")

	// сформировать access_token
	ts := time.Now()
	claims := jwt.MapClaims{}
	claims["iss"] = "perimeter"
	claims["sub"] = "access_token"
	claims["appId"] = request.AppId
	claims["user"] = request.AppId
	claims["username"] = "Test user FullName"
	claims["password"] = "Test user FullName"
	claims["access_token"] = "gthbvtnhulc64jrtq"
	claims["refresh_token"] = "gthbvtnhulc64jrtq"
	claims["exp"] = ts.Add(time.Minute * config.AccessTokenLifeMinutes).Unix()
	claims["ts"] = ts
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// отдаем токен клиенту
	tokenString, err := token.SignedString(request.SecretSign())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Println(tokenString)
	}

	rw.Write([]byte(tokenString))
}
