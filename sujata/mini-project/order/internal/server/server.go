package server

import (
	"context"
	"net/http"
	"order/config"
	mongodao "order/internal/dao"
	"order/internal/services"
	"order/util"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	dao, err := intializeDao(webServerConfig)
	if err != nil {
		return err
	}

	// Initialize services
	services.InitCreateOrderService(&routerConfigs, dao)
	services.InitGetOrderService(&routerConfigs, dao)
	services.InitSetOrderStatusService(&routerConfigs, dao)

	server := NewServer(webServerConfig)
	server.Router.InitializeRouter(&routerConfigs)

	log.Info("Server starting on PORT: ", webServerConfig.Port)
	err = http.ListenAndServe(":"+webServerConfig.Port, *server.Router)
	if err != nil {
		return err
	}

	return nil
}

func intializeDao(config *config.WebServerConfig) (mongodao.MongoDAO, error) {
	// Initialize mongo database connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.WithError(err).Error("an error occurred while initializing mongo connection")
		return nil, err
	}

	dao := mongodao.InitMongoDAO(client, config)
	return dao, nil
}
