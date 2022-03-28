package server

import (
	"cart/internal/literals"
	"cart/internal/middleware"
	"cart/internal/server/handlers"
	"cart/util"
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

	s.HandleFunc(literals.CartProductEndpoint,
		middleware.AuthHandler(handlers.AddProductToCartHandler(routerConfig))).
		Methods(http.MethodPost).
		Name(literals.AddProductToCartAPIName)

	s.HandleFunc(literals.CartProductEndpoint,
		middleware.AuthHandler(handlers.DeleteProductFromCartHandler(routerConfig))).
		Methods(http.MethodDelete).
		Name(literals.DeleteProductFromCartAPIName)

	s.HandleFunc(literals.CartProductEndpoint,
		middleware.AuthHandler(handlers.GetCartHandler(routerConfig))).
		Methods(http.MethodGet).
		Name(literals.GetCartAPIName)
}
