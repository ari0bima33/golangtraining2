package routes

import (
	"net/http"

	"final-project/domain/general"
	"final-project/handlers/core"

	"github.com/gorilla/mux"
)

func photo(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	router.HandleFunc("/photos", handler.Photo.Photo.CreatePhoto).Methods(http.MethodPost)
	router.HandleFunc("/photos", handler.Photo.Photo.GetList).Methods(http.MethodGet)
	router.HandleFunc("/photos", handler.Photo.Photo.UpdatePhoto).Methods(http.MethodPut)
	router.HandleFunc("/photos", handler.Photo.Photo.DeletePhoto).Methods(http.MethodDelete)

}
