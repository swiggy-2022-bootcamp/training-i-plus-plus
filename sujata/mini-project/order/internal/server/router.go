package server

import (
	"net/http"
	"order/internal/literals"
	"order/internal/middleware"
	"order/internal/server/handlers"
	"order/util"

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

// @title Orders API
// @version 1.0
// @description This is a sample serice for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (r *Router) InitializeRoutes(routerConfig *util.RouterConfig) {
	s := (*r).PathPrefix(routerConfig.WebServerConfig.RoutePrefix).Subrouter()

	// CreateOrder godoc
	// @Summary Create a new order
	// @Description Create a new order with the input paylod
	// @Tags orders
	// @Accept  json
	// @Produce  json
	// @Param order body Order true "Create order"
	// @Success 200 {object} Order
	// @Router /orders [post]
	s.HandleFunc(literals.CreateOrderEndpoint,
		middleware.AuthHandler(handlers.CreateOrderHandler(routerConfig))).
		Methods(http.MethodPost).
		Name(literals.CreateOrderAPIName)

	// GetOrders godoc
	// @Summary Get details of all orders
	// @Description Get details of all orders
	// @Tags orders
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} Order
	// @Router /orders [get]
	s.HandleFunc(literals.GetOrderEndpoint,
		middleware.AuthHandler(handlers.GetOrderStatusHandler(routerConfig))).
		Methods(http.MethodGet).
		Name(literals.GetOrderAPIName)

	s.HandleFunc(literals.SetOrderStatusEndpoint,
		middleware.AuthHandler(handlers.SetOrderStatusHandler(routerConfig))).
		Methods(http.MethodPost).
		Name(literals.SetOrderStatusAPIName)
}
