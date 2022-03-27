package server

import (
	"net/http"
	"search/internal/literals"
	"search/internal/server/handlers"
	"search/util"

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

	s.HandleFunc(literals.SearchProductEndpoint,
		handlers.SearchProductHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.SearchProductAPIName)
}
