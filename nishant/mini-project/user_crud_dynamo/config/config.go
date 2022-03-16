package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DB struct {
		Endpoint string
		Region   string
		ID       string
		Secret   string
	}
}

func loadConfig() Config {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatalln("error opening file : ", err)
	}
	defer file.Close()
	log.Println("Reading Config file ", file)
	decoder := json.NewDecoder(file)
	c := Config{}
	if err := decoder.Decode(&c); err != nil {
		log.Fatalln("error:", err)
	}
	return c
}

var c = loadConfig()

func C() Config {
	return c
}
