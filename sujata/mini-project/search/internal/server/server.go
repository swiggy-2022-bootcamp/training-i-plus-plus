package server

import (
	"net/http"
	"search/config"
	"search/internal/services"
	"search/util"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Configuration *config.WebServerConfig
	Router        *Router
}

// NewServer creates the new server and sets the server configurations.
func NewServer(config *config.WebServerConfig) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
	}
	return server
}

// RunServer initializes the server along with all the microservice dependencies.
// It starts the server and returns nil as error if server starts successfully otherwise
// returns the error.
func RunServer() error {
	webServerConfig, err := config.FromEnv()
	if err != nil {
		return err
	}

	routerConfigs := util.RouterConfig{
		WebServerConfig: webServerConfig,
	}

	// Initialize services
	services.InitSearchProductService(&routerConfigs)

	server := NewServer(webServerConfig)
	server.Router.InitializeRouter(&routerConfigs)

	log.Info("Server starting on PORT: ", webServerConfig.Port)
	err = http.ListenAndServe(":"+webServerConfig.Port, *server.Router)
	if err != nil {
		return err
	}

	return nil
}
