package server

import (
	"auth/internal/literals"
	"auth/internal/server/handlers"
	"auth/util"
	"net/http"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRouter(routerConfig *util.RouterConfig) {
	r.InitializeRoutes(routerConfig)
}

func (r *Router) InitializeRoutes(routerConfig *util.RouterConfig) {
	s := (*r).PathPrefix(routerConfig.WebServerConfig.RoutePrefix).Subrouter()
	// Swagger
	s.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	s.HandleFunc(literals.SignupEndpoint,
		handlers.SignupHandler(routerConfig)).
		Methods(http.MethodPost).
		Name(literals.SignupAPIName)

	s.HandleFunc(literals.SigninEndpoint,
		handlers.SigninHandler(routerConfig)).
		Methods(http.MethodPost).
		Name(literals.SigninAPIName)
}
