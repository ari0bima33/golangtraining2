package routes

import (
	"net/http"

	"final-project/domain/general"
	"final-project/handlers/core"

	"github.com/gorilla/mux"
)

func socialmedia(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/socialmedias", handler.Socialmedia.Socialmedia.CreateSocialmedia).Methods(http.MethodPost)
	router.HandleFunc("/socialmedias", handler.Socialmedia.Socialmedia.GetList).Methods(http.MethodGet)
	router.HandleFunc("/socialmedias", handler.Socialmedia.Socialmedia.UpdateSocialmedia).Methods(http.MethodPut)
	router.HandleFunc("/socialmedias", handler.Socialmedia.Socialmedia.DeleteSocialmedia).Methods(http.MethodDelete)

}
