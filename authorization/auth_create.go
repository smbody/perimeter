package authorization

import (
	"fmt"
	"net/http"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/data"
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

	username, password, ok := req.BasicAuth()
	if !ok {
		http.Error(rw, config.HttpErrorBadRequestMethod, http.StatusBadRequest)
		return
	}

	user, err_u := data.CreateUser(request.AppId, username, password)
	if err_u != nil {
		fmt.Println(err_u.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Println("create user id=", user.Id)

	// создать токены
	tokens := data.Login(request.AppId, user.Id)

	fmt.Println("create token")
	fmt.Println("access token", tokens.Access)
	// сформировать токен в ответ
	token := authResponse(request.AppId, user, tokens)

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
