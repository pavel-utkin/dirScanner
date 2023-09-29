package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type ConfigClient struct {
	GRPC string `env:"AddressGRPC"`
}

// NewConfigClient - creates a new instance with the configuration for the client
func NewConfig() *ConfigClient {
	configClient := ConfigClient{
		GRPC: "localhost:8080", // default value
	}
	flag.StringVar(&configClient.GRPC, "g", configClient.GRPC, "Server address")
	flag.Parse()
	err := env.Parse(&configClient)
	if err != nil {
		log.Fatal(err)
	}
	return &configClient
}
