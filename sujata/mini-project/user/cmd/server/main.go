package main

import (
	"user/internal/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := server.RunServer(); err != nil {
		log.WithField("Error: ", err).Fatalf("Server quitting...")
	}
}
