package main

import (
	"fmt"
	"net/http"

	"github.com/smbody/perimeter/authorization"
	"github.com/smbody/perimeter/config"
)

func main() {
	fmt.Println("Perimeter activate")

	authorization.RegisterHandlers()

	port := config.Port
	fmt.Println("Start on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
