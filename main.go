package main

import (
	"fmt"
	"net/http"

	"github.com/smbody/perimeter/authorization"
	"github.com/smbody/perimeter/config"
	"github.com/smbody/perimeter/dao"
)

func main() {
	fmt.Println("Perimeter activate")

	authorization.RegisterHandlers()
	dao.Register(dao.Mongo())

	port := config.Port
	fmt.Println("Start on port", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
