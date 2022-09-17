package main

import (
	"fmt"
	"log"

	"github.com/edmarfelipe/go-prometheus/config"
	"github.com/edmarfelipe/go-prometheus/server"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("could not load config: %w", err))
	}

	server := server.New(config)
	server.Run(":8081")
}
