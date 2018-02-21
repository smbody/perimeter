package authorization

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/auth/create", authCreate)
	http.HandleFunc("/auth/login", authLogin)
}
