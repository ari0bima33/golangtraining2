package main

import (
	"fmt"
	"net/http"

	"final-project/cmd/core/config"
	"final-project/cmd/core/routes"

	"github.com/gorilla/handlers"
)

func main() {
	conf, err := config.GetCoreConfig()
	if err != nil {
		panic(err)
	}

	handler, log, err := config.NewRepoContext(conf)
	if err != nil {
		panic(err)
	}

	headers := handlers.AllowedHeaders(conf.Route.Headers)
	methods := handlers.AllowedMethods(conf.Route.Methods)
	origins := handlers.AllowedOrigins([]string{conf.Route.Origins.InternalTools})
	credentials := handlers.AllowCredentials()

	router := routes.GetCoreEndpoint(conf, handler, log)

	port := fmt.Sprintf(":%s", conf.App.Port)
	log.Info("server listen to port ", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins, credentials)(router)))
}
