package routes

import (
	"net/http"

	"final-project/domain/general"
	"final-project/handlers/core"

	"github.com/gorilla/mux"
)

func user(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/user/register", handler.User.User.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/user/login", handler.User.User.UserLogin).Methods(http.MethodPost)
	routerJWT.HandleFunc("/users", handler.User.User.UpdateUser).Methods(http.MethodPut)
	routerJWT.HandleFunc("/users", handler.User.User.DeleteUser).Methods(http.MethodDelete)
}
