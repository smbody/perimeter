package authorization

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/data"
)

func RegisterHandlers() {
	http.HandleFunc("/auth", authUser)
	http.HandleFunc("/auth/create", authCreate)
	http.HandleFunc("/auth/login", authLogin)
	http.HandleFunc("/auth/token", authToken)
}

type AuthRequest struct {
	AppId  string
	secret string
	Token  *jwt.Token
	Claims jwt.MapClaims
}

func (req *AuthRequest) Secret() string {
	if req.secret == "" {
		req.secret = config.GetSecretApp(req.AppId)
	}

	return req.secret
}

func (req *AuthRequest) SecretSign() []byte {
	return []byte(req.Secret())
}

func authRequest(req *http.Request) (*AuthRequest, string) {
	if req.Method != http.MethodPost {
		return nil, config.HttpErrorBadRequestMethod
	}

	request := new(AuthRequest)

	// для начала определим app
	appId := req.URL.Query().Get("appId")
	if appId != "" {
		request.AppId = appId
	} else {
		return nil, config.HttpErrorBadRequestAppId
	}

	if request.Secret() == "" {
		return nil, config.HttpErrorBadRequestApp
	}

	// выцепим токен
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return nil, err.Error()
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(string(body), claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(request.Secret()), nil
		})

	if err != nil || !token.Valid {
		return nil, err.Error()
	} else {
		request.Token = token
		request.Claims = claims
	}

	return request, ""
}

func authResponse(appId string, user *data.User) *jwt.Token {

	// сформировать токен
	ts := time.Now()
	claims := jwt.MapClaims{}
	claims["iss"] = "perimeter"
	claims["sub"] = "access_token"
	claims["appId"] = appId
	claims["user"] = user.Id
	claims["userName"] = user.FullName()
	claims["access_token"] = user.AccessToken(appId)
	claims["refresh_token"] = user.RefreshToken(appId)
	claims["exp"] = ts.Add(time.Minute * config.AccessTokenLifeMinutes).Unix()
	claims["ts"] = ts
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}
