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
	if len(header) == 2 {
		user, err := data.ValidateToken(appId, header[1])
		if user != nil {
			// если удалось аутентифицировать, отдаем юзера
			rw.Header().Set("Content-Type", "app/json")
			rw.Write(user.Json())
			return
		} else {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Error: ", config.HttpErrorBadRequestToken)
	}

	// если не удалось - вернуть 401
	http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
