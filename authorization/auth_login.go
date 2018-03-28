package authorization

import (
	"fmt"
	"net/http"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/data"
)

func authLogin(rw http.ResponseWriter, req *http.Request) {

	request, err := authRequest(req)
	if err != "" {
		fmt.Println(err)
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// получить user-a
	username := request.Claims["username"]
	password := request.Claims["password"]
	if username == nil || password == nil {
		fmt.Println(config.AuthErrorBadUserName)
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	user, err_u := data.GetUser(request.AppId, username.(string), password.(string))
	if err_u != nil {
		fmt.Println(err_u.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// создать токены
	tokens := data.Login(request.AppId, user.Id)

	// сформировать токен в ответ
	token := authResponse(request.AppId, user, tokens)

	// отдаем токен клиенту
	tokenString, err_t := token.SignedString(request.SecretSign())
	if err_t != nil {
		fmt.Println(err_t.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	rw.Write([]byte(tokenString))
}
