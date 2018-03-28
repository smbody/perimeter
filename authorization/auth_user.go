package authorization

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/data"
)

func authUser(rw http.ResponseWriter, req *http.Request) {
	// GET only
	if req.Method != http.MethodGet {
		fmt.Println(config.HttpErrorBadRequestMethod)
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// для начала определим app
	appId := req.URL.Query().Get("appId")

	// токен
	header := strings.Split(req.Header.Get("Authorization"), " ")
	if len(header) != 2 {
		fmt.Println("Error: ", config.HttpErrorBadRequestToken)
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}

	// проверим токен
	token, err := data.ValidateToken(appId, header[1])
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// информация по пользователю
	user, err_u := data.GetUserById(appId, token.UserId)
	if err_u != nil {
		fmt.Println(err_u.Error())
		http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// если все удалось, отдаем юзера
	rw.Header().Set("Content-Type", "app/json")
	rw.Write(user.Json())
}
