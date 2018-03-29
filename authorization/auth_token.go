package authorization

import (
	"fmt"
	"net/http"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/data"
)

func authToken(rw http.ResponseWriter, req *http.Request) {

	request, err := authRequest(req)
	if err != "" {
		fmt.Println(err)
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// для начала проверим app
	appId := request.Claims["appId"]
	if request.AppId != appId {
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// проверим токен
	refresh_token := request.Claims["refresh_token"]
	if refresh_token == nil {
		fmt.Println(config.AuthErrorBadRefreshToken)
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// обновляем токены
	tokens, err_u := data.RefreshToken(request.AppId, refresh_token.(string))
	if err_u != nil {
		fmt.Println(err_u.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// информация по пользователю
	user, err_u := data.GetUserById(request.AppId, tokens.UserId)
	if err_u != nil {
		fmt.Println(err_u.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	fmt.Println("refresh token")
	fmt.Println("access token", tokens.Access)
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
