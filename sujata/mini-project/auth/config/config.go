package config

import (
	"auth/internal/literals"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type WebServerConfig struct {
	Port         string `required:"true" split_words:"true"`
	RoutePrefix  string `required:"false" split_words:"true" default:"/auth"`
	Db           string `required:"true" split_words:"true"`
	DbCollection string `required:"true" split_words:"true"`
	MongoUrl     string `required:"true" split_words:"true"`
}

func FromEnv() (*WebServerConfig, error) {
	cfgFilename := "../../etc/config.localhost.env"
	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Println("No config files found to load to env.")
	}

	webServerConfig := &WebServerConfig{}
	err := envconfig.Process(literals.AppPrefix, webServerConfig)
	if err != nil {
		return nil, err
	}

	return webServerConfig, nil
}
