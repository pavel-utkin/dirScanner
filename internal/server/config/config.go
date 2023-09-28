package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	AddressGRPC string `env:"AddressGRPC"`
}

// NewConfig - creates a new instance with the configuration for the server
func NewConfig() *Config {
	// Set default values
	configServer := Config{
		AddressGRPC: "localhost:8080",
	}

	flag.StringVar(&configServer.AddressGRPC, "g", configServer.AddressGRPC, "Server address GRPC")
	flag.Parse()
	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}
	return &configServer
}
