package main

import (
	"dirScanner/internal/server/api"
	grpchandler "dirScanner/internal/server/api/grpc"
	"dirScanner/internal/server/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	serverConfig := config.NewConfig()
	handlerGrpc := grpchandler.NewHandler(logger)
	go api.StartGRPCService(handlerGrpc, serverConfig)
}
