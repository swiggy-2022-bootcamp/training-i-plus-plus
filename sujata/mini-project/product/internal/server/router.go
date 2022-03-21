package server

import (
	"net/http"
	"product/internal/literals"
	"product/internal/middleware"
	"product/internal/server/handlers"
	"product/util"

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

	s.HandleFunc(literals.AddProductEndpoint,
		middleware.AuthHandler(handlers.AddProductHandler(routerConfig))).
		Methods(http.MethodPost).
		Name(literals.AddProductAPIName)
}
