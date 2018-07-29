package main

import (
	"goproject/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//handler-->validation(1.request  2.user)-->business logic-->respons
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUser)
	router.POST("/user/:user_name", handler.Login)
	return router

}

func main() {
	http.ListenAndServe(":8000", RegisterHandlers())
}
