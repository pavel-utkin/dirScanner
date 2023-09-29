package main

import (
	"context"
	"dirScanner/internal/server/api"
	grpchandler "dirScanner/internal/server/api/grpc"
	"dirScanner/internal/server/config"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	logger := logrus.New()
	serverConfig := config.NewConfig()
	handlerGrpc := grpchandler.NewHandler(logger)

	ctx, cnl := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cnl()

	go api.StartGRPCService(handlerGrpc, serverConfig)

	<-ctx.Done()
	logger.Info("server shutdown on signal with:", ctx.Err())
}
