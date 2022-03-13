package server

import (
	"auth/internal/literals"
	"auth/internal/server/handlers"
	"auth/util"
	"net/http"

	"github.com/gorilla/mux"
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

	s.HandleFunc(literals.SignupEndpoint,
		handlers.SignupHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.SignupAPIName)
}
