package routes

import (
	"net/http"

	"final-project/domain/general"
	"final-project/handlers/core"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetCoreEndpoint(conf *general.SectionService, handler core.Handler, log *logrus.Logger) *mux.Router {
	parentRoute := mux.NewRouter()

	jwtRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()
	nonJWTRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()
	publicRoute := parentRoute.PathPrefix(conf.App.Endpoint).Subrouter()

	// Renew Access Token Endpoint.
	publicRoute.HandleFunc("/renew-token", handler.Token.RenewAccessToken).Methods(http.MethodGet)

	// Middleware for public API
	nonJWTRoute.Use(handler.Public.AuthValidator)

	// Middleware
	if conf.Authorization.JWT.IsActive {
		log.Info("JWT token is active")
		jwtRoute.Use(handler.Token.JWTValidator)
	}

	// Get Endpoint.
	user(publicRoute, jwtRoute, conf, handler)
	photo(jwtRoute, jwtRoute, conf, handler)
	comment(jwtRoute, jwtRoute, conf, handler)
	socialmedia(jwtRoute, jwtRoute, conf, handler)
	return parentRoute
}
