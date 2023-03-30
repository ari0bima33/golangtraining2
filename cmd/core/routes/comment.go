package routes

import (
	"net/http"

	"final-project/domain/general"
	"final-project/handlers/core"

	"github.com/gorilla/mux"
)

func comment(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/comments", handler.Comment.Comment.CreateComment).Methods(http.MethodPost)
	router.HandleFunc("/comments", handler.Comment.Comment.GetList).Methods(http.MethodGet)
	router.HandleFunc("/comments", handler.Comment.Comment.UpdateComment).Methods(http.MethodPut)
	router.HandleFunc("/comments", handler.Comment.Comment.DeleteComment).Methods(http.MethodDelete)

}
